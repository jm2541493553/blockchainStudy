package blc

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

// 代表256位hash前面至少有16个0
const targetBit = 16

type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数据存储
}

// 验证hash是否有效
func (pow *ProofOfWork) IsValid() bool {
	// pow.Block.Hash 和 pwo.target作比较
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PreBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(pow.Block.Height),
		},
		[]byte{},
	)
	return data
}
func (pow *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	var hashInt big.Int //存储新生成的hash
	var hash [32]byte
	for {
		//1.将所有的block拼接转化成字节数组
		dataBytes := pow.prepareData(nonce)
		//2.生成hash
		hash = sha256.Sum256(dataBytes)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//3.判断hash的有效性，死循环，满足条件，跳出循环
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1
	}
	return hash[:], int64(nonce)
}

// NewProofOfWork 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	// 1. big.Int对象 1
	// 0000 0001 至少有两个0难度，左移6位0100 0000==》64
	// 但凡生成的hash前两个为0，小于target

	//初始值为1的target
	target := big.NewInt(1)
	//左移256-targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}

// 判断当前区块是否有效
