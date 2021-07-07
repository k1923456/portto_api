package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	Num        uint
	Hash       string
	Time       uint
	ParentHash string
}

type Transaction struct {
	gorm.Model
	Hash  string
	From  string
	To    string
	Nonce uint
	Data  string
	Value string
	Logs  string
}

type BlockTransaction struct {
	gorm.Model
	Num  uint
	Hash string
}

var wg sync.WaitGroup

func insertDB(ctx context.Context, client *ethclient.Client, db *gorm.DB, chainId *big.Int, startBLock int, endBlock int) {
	defer wg.Done()
	retry := 5
	for i := startBLock; i <= endBlock; i++ {
		// TODO: Here would fail with server graceful down
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			fmt.Printf("get block failed, retry = %v, block number us %v, err is %v\n", retry, i, err)
			for retry > 0 {
				time.Sleep(30000 * time.Millisecond) // 30 sec
				retry--
				i--
				continue
			}
			return
		}
		db.Create(&Block{
			Num:        uint(block.Number().Uint64()),
			Hash:       block.Hash().String(),
			Time:       uint(block.Header().Time),
			ParentHash: block.ParentHash().String(),
		})
		transactions := block.Transactions()
		for i := 0; i < transactions.Len(); i++ {
			msg, err := transactions[i].AsMessage(types.NewEIP155Signer(chainId), nil)
			if err != nil {
				fmt.Printf("get Message failed, err is %v\n", err)
			}
			receipt, err := client.TransactionReceipt(ctx, transactions[i].Hash())
			if err != nil {
				fmt.Printf("get Receipt failed, err is %v\n", err)
			}
			logs, _ := json.Marshal(receipt.Logs)
			to := msg.To()
			toString := ""
			if to != nil {
				toString = to.String()
			}
			db.Create(&BlockTransaction{
				Num:  uint(block.Number().Uint64()),
				Hash: transactions[i].Hash().String(),
			})
			db.Create(&Transaction{
				Hash:  transactions[i].Hash().String(),
				From:  msg.From().String(),
				To:    toString,
				Nonce: uint(msg.Nonce()),
				Data:  hex.EncodeToString(msg.Data()),
				Value: msg.Value().String(),
				Logs:  string(logs),
			})
		}
	}
}

func main() {
	// start := time.Now()
	numWindow := 1000
	maxBlockNum := 10000

	// Connect to DB
	dsn := "host=localhost user=user password=password dbname=user port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to get DB interface")
	}

	// Set connection pool
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	// Migrate the schema
	db.AutoMigrate(&Block{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&BlockTransaction{})

	url := "https://data-seed-prebsc-1-s1.binance.org:8545/"
	ctx := context.Background()
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	chainId, err := client.NetworkID(ctx)
	if err != nil {
		fmt.Printf("get Network ID failed, err is %v\n", err)
	}
	// maxBlockNum, err := client.BlockNumber()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	initialCount := 0
	for {
		if initialCount < maxBlockNum {
			for i := initialCount; i < int(maxBlockNum); i += numWindow {
				go insertDB(ctx, client, db, chainId, i, i+numWindow-1)
				wg.Add(1)
			}
			wg.Wait()
			initialCount += maxBlockNum
		}
		// elapsed := time.Since(start)
		// fmt.Printf("Run time took %s\n", elapsed)
		fmt.Printf("Scan done, Got %v blocks", initialCount)
		time.Sleep(3000 * time.Millisecond)
	}
}
