package main

import (
	"bytes"
	"fmt"
)

type MemoryStore struct {
	lastHash []byte
	data     map[string]*Block
}

func (m *MemoryStore) Load(hash []byte) (*Block, error) {
	x := fmt.Sprintf("%x", hash)

	if _, ok := m.data[x]; ok {
		return m.data[x], nil
	}

	return nil, fmt.Errorf("block not found")

}

func (m *MemoryStore) LastHash() ([]byte, error) {

	if len(m.lastHash) == 0 {
		return nil, ErrorNotInitialize
	}

	return m.lastHash, nil

}

func (m *MemoryStore) Append(block *Block) error {
	if !bytes.Equal(block.PrevHash, m.lastHash) {
		return fmt.Errorf("new block prevHash not Equal with prev block hash")
	}
	x := fmt.Sprintf("%x", block.Hash)
	if _, ok := m.data[x]; ok {
		return fmt.Errorf("duplicate hash ...")
	}

	m.data[x] = block
	m.lastHash = block.Hash

	return nil

}

func createNewMemoryStore() Store {
	return &MemoryStore{
		data : make(map[string]*Block),
	}
}
