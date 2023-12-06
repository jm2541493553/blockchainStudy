package main

import (
	"blockchainStudy/basic_prototype/blc"
	"fmt"
)

func main() {
	blockChain := blc.CreateBlockChainWithGenesisBlock()
	// 新区块
	blockChain.AddBlockToBlockChain("Send 100RMB to rws", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 200RMB to rws", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 300RMB to rws", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 400RMB to rws", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 500RMB to rws", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	for _, block := range blockChain.Blocks {
		fmt.Println(block)
	}
}
