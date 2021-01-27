package main

import (
	"bytes"
	"fmt"
)

type Blockchain struct {
	Blocks     []*Block
	Mask       []byte
	Difficulty int
}

func (bc *Blockchain) String() string {
	var ret string
	for _, v := range bc.Blocks {
		ret += v.String()
	}
	return ret
}

func (bc *Blockchain) Add(data string) {
	lenBlockchain := len(bc.Blocks)

	if lenBlockchain <= 0 {
		panic("First Create Blockchain then add block !")
	}
	prevHash := bc.Blocks[lenBlockchain-1].Hash

	bc.Blocks = append(bc.Blocks, NewBlock(data, bc.Mask, prevHash))

}

func NewBlockchain(difficulty int) *Blockchain {
	mask := GenerateMask(difficulty)
	bc := Blockchain{
		Difficulty: difficulty,
		Mask:       mask,
		Blocks:     []*Block{NewBlock("Genesis Block", mask, []byte{})},
	}

	return &bc
}

func (bc *Blockchain) Validate() error {
	for i, v := range bc.Blocks {
		if err := v.Validate(bc.Mask); err != nil {
			return fmt.Errorf("Blockchain is not valid : \n %w", err)
		}
		if i == 0 {
			continue
		}
		if !bytes.Equal(v.PrevHash, bc.Blocks[i-1].Hash) {
			return fmt.Errorf("the order of blocks is invalid in Block : %d,\n  it is:\n%x \nshould be :\n %x .",
				i, v.PrevHash, bc.Blocks[i-1].Hash)
		}
	}

	return nil

}
