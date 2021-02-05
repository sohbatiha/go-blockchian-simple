package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type configData struct {
	LastHash []byte
}

type FileStore struct {
	config     *configData
	path       string
	configPath string
}

func (f *FileStore) Load(hash []byte) (*Block, error) {

	blockPath := filepath.Join(f.path, fmt.Sprintf("%x.json", hash))
	var block Block

	if err := readJson(blockPath, &block); err != nil {
		return nil, fmt.Errorf("block with this hash not found , \n path : %s ", blockPath)
	}

	return &block, nil
}

func (f *FileStore) LastHash() ([]byte, error) {
	if len(f.config.LastHash) == 0 {
		return nil, ErrorNotInitialize
	}
	return f.config.LastHash, nil
}

func (f *FileStore) Append(block *Block) error {
	if !bytes.Equal(block.PrevHash, f.config.LastHash) {
		return fmt.Errorf("block prevHash not Equal with last block hash")
	}

	blockPath := filepath.Join(f.path, fmt.Sprintf("%x.json", block.Hash))

	if err := writeJson(blockPath, block); err != nil {
		return err
	}

	f.config.LastHash = block.Hash

	if err := writeJson(f.configPath, f.config); err != nil {
		return err
	}

	return nil
}

func CreateNewFileStore(root string) Store {
	s := &FileStore{
		path:       root,
		config:     &configData{},
		configPath: filepath.Join(root, "config.json"),
	}

	if err := readJson(s.configPath, s.config); err != nil {
		fmt.Println("Empty store ...")
		s.config.LastHash = nil
	}

	return s
}

func readJson(path string, v interface{}) error {
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()

	dec := json.NewDecoder(fl)

	if err := dec.Decode(v); err != nil {
		return err
	}

	return nil
}

func writeJson(path string, v interface{}) error {
	fl, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fl.Close()

	enc := json.NewEncoder(fl)
	enc.SetIndent("", "   ")

	if err := enc.Encode(v); err != nil {
		return err
	}

	return nil
}
