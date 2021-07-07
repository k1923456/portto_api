package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

type APIBlock struct {
	Num        uint   `json:"block_num"`
	Hash       string `json:"block_hash"`
	Time       uint   `json:"block_time"`
	ParentHash string `json:"parent_hash"`
}

func GetBlocks(c *gin.Context) {
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	MaxBLockNum := models.GetMaxBlockNum()
	var res = make([]APIBlock, limit)
	for i := 0; i < int(limit); i++ {
		index := MaxBLockNum - i
		tmp := models.GetBlock(uint(index))
		res[i].Num = tmp.Num
		res[i].Hash = tmp.Hash
		res[i].Time = tmp.Time
		res[i].ParentHash = tmp.ParentHash
	}
	fmt.Println(res)
	c.JSON(200, res)
}

func GetBlock(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	block := models.GetBlock(uint(id))
	transactions := models.GetTransactionsByBlockNum(uint(id))
	c.JSON(200, gin.H{
		"block_num":    block.Num,
		"block_hash":   block.Hash,
		"block_time":   block.Time,
		"parent_hash":  block.ParentHash,
		"transactions": transactions,
	})
}

func GetTransaction(c *gin.Context) {
	hash := c.Param("txHash")
	transaction := models.GetTransaction(hash)
	logs, _ := json.Marshal([]byte(transaction.Logs))
	c.JSON(200, gin.H{
		"tx_hash": transaction.Hash,
		"from":    transaction.From,
		"to":      transaction.To,
		"nonce":   transaction.Nonce,
		"data":    transaction.Data,
		"value":   transaction.Value,
		"logs":    logs,
	})
}
