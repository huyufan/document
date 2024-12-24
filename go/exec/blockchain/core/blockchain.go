package core

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	Block := GenerateGenesisBlock()
	bc := new(BlockChain)
	bc.AppendBlock(Block)
	return bc
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	nextBlock := GenerateNewBlock(preBlock, data)
	bc.AppendBlock(nextBlock)
}

func (bc *BlockChain) AppendBlock(block *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, block)
	} else {
		if bc.isValid(block) {
			bc.Blocks = append(bc.Blocks, block)
		}
	}
}

func (bc *BlockChain) isValid(block *Block) bool {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	if prevBlock.Index != block.Index-1 {
		return false
	}
	if prevBlock.Hash != block.PrevBlockHash {
		return false
	}
	return true
}
