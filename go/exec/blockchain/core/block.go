package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  `json:"index"`     //区块编号
	Data          string `json:"data"`      //区块保持数据
	Hash          string `json:"hash"`      //当前区块的Hash值
	Timestamp     int64  `json:"timestamp"` // 时间戳
	PrevBlockHash string `json:"prevHash"`  // 上一个区块的hash值
}

func GenerateNewBlock(prevBlock *Block, data string) *Block {
	newBlock := new(Block)
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func calculateHash(block *Block) string {
	blockHash := string(block.Index) + string(block.Timestamp) + block.Hash + block.Data + block.PrevBlockHash
	blockBytes := sha256.Sum256([]byte(blockHash))
	return hex.EncodeToString(blockBytes[:])

}

func GenerateGenesisBlock() *Block {
	block := new(Block)
	block.Index = -1
	block.Timestamp = time.Now().Unix()
	block.Hash = ""
	return GenerateNewBlock(block, "Genesis Block")
}
