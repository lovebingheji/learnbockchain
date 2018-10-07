package BLC

import (
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64

	//2.上一个区块的Hash
	PrevBlockHash []byte

	//3.交易数据
	Data []byte

	//4.时间戳
	Timestamp int64

	//5.本区块的Hash
	Hash []byte
}

//创建新的区块结构
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	return &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
}
