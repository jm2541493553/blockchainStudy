package blc

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

// SetHash 设置hash
func (block *Block) SetHash() {
	//1.height - >[]byte
	heightBytes := IntToHex(block.Height)
	//fmt.Println(heightBytes)
	//2.timestamp->[]byte
	//2~36进制
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	//fmt.Println(timeBytes)
	//3.拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PreBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
	//4.生成hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
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
	// 设置Hash
	block.SetHash()
	return block
}

// CreateGenesisBlock 创建创世块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
