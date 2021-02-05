package main

import (
	"errors"
	"fmt"
)

var ErrorNotInitialize = errors.New("blockchain not initialize")

type Blockchain struct {
	Mask       []byte
	Difficulty int
	store      Store
}

func (bc *Blockchain) Print() error {

	return Iterate(bc.store, func(block *Block) error {
		fmt.Println(block)
		return nil
	})

}

func (bc *Blockchain) Add(data string) (*Block, error) {
	lastHash, err := bc.store.LastHash()
	if err != nil {
		return nil, err
	}
	block := NewBlock(data, bc.Mask, lastHash)

	if err := bc.store.Append(block); err != nil {
		return nil, err
	}

	return block, nil

}

func NewBlockchain(difficulty int, store Store) (*Blockchain, error) {
	mask := GenerateMask(difficulty)
	bc := Blockchain{
		Difficulty: difficulty,
		Mask:       mask,
		store:      store,
	}

	_, err := store.LastHash()

	if err == nil {
		return &bc, nil
	}

	if !errors.Is(err, ErrorNotInitialize) {
		return nil, err
	}

	if errors.Is(err, ErrorNotInitialize) {
		err := store.Append(NewBlock("Genesis Block", mask, []byte{}))
		if err != nil {
			return nil, err
		}
	}

	return &bc, nil
}

func (bc *Blockchain) Validate() error {

	return Iterate(bc.store, func(block *Block) error {
		if err := block.Validate(bc.Mask); err != nil {
			return fmt.Errorf("Blockchain is not valid : \n %w", err)
		}
		return nil
	})

}
