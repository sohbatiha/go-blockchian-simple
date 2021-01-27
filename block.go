package main

import (
	"bytes"
	"fmt"
	"time"
)

type Block struct {
	Time time.Time
	Data []byte

	Hash     []byte
	PrevHash []byte

	Nonce int32
}

func (b *Block) Validate(mask []byte) error {
	hash := GenerateHash(b.Time.UnixNano(), b.Data, b.PrevHash, b.Nonce)
	if !bytes.Equal(hash, b.Hash) {
		return fmt.Errorf("Hash is invalid \nit should be :\n%x \nbut is :\n%x", b.Hash, hash)
	}
	if !checkHashCondition(mask, b.Hash) {
		return fmt.Errorf("hash is diffirent than difficulty condition .\nmask: %x ,\nhash: %x\n", mask, b.Hash)
	}
	return nil
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"\n-------------\nTime : %s \nData : %s \nHash : %x \nPrevHash : %x\n",
		b.Time, b.Data, b.Hash, b.PrevHash)
}

func NewBlock(data string, mask []byte, prevHash []byte) *Block {
	b := Block{
		Data:     []byte(data),
		Time:     time.Now(),
		PrevHash: prevHash,
	}
	b.Hash, b.Nonce = GenerateHashWithDifficulty(mask, b.Time.UnixNano(), b.Data, b.PrevHash)

	return &b

}
