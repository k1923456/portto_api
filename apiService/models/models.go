package models

import (
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

var db *gorm.DB

func Init(_db *gorm.DB) {
	db = _db
}

func GetBlock(blockNum uint) (block Block) {
	db.First(&block, "num = ?", blockNum)
	return block
}

func GetTransaction(hash string) (transaction Transaction) {
	db.First(&transaction, "hash = ?", hash)
	return transaction
}

func GetTransactionsByBlockNum(blockNum uint) (transactions []string) {
	var res []BlockTransaction
	db.Raw("SELECT hash FROM block_transactions WHERE num = ?", blockNum).Scan(&res)
	for _, v := range res {
		transactions = append(transactions, v.Hash)
	}
	return transactions
}

func GetMaxBlockNum() (blockNum int) {
	type MaxStruct struct {
		Max int `json:"max"`
	}
	tmp := MaxStruct{}
	db.Raw("SELECT max(block_transactions.num) AS max FROM block_transactions").Scan(&tmp)
	return int(tmp.Max)
}
