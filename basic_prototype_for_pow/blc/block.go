package blc

import (
	"time"
)

type Block struct {
	// 1.区块高度
	Height int64
	// 2.上一个区块的hash
	PreBlockHash []byte
	// 3.交易数据
	Data []byte
	// 4，时间戳
	Timestamp int64
	// 5.Hash
	Hash []byte
	// 6.Nonce
	Nonce int64
}

// NewBlock 1.创建新的区块
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	// 创建区块
	block := &Block{
		Height:       height,
		PreBlockHash: preBlockHash,
		Data:         []byte(data),
		Timestamp:    time.Now().Unix(),
		Hash:         nil,
	}
	// 调用工作量证明方法返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	// 000000开头
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// CreateGenesisBlock 创建创世块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
