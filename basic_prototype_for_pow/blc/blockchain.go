package blc

type BlockChain struct {
	Blocks []*Block //存储有效的区块
}

// AddBlockToBlockChain 添加区块到区块链
func (blockchain *BlockChain) AddBlockToBlockChain(data string, height int64, preBlockHash []byte) {
	newBlock := NewBlock(data, height, preBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// CreateBlockChainWithGenesisBlock 创建带有创世区块的区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Block.....")
	return &BlockChain{
		Blocks: []*Block{genesisBlock},
	}
}
