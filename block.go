package main

import (
	"bytes"
	"fmt"
	"time"
)

type Block struct {
	Time     time.Time
	Data     []byte
	Hash     []byte
	PrevHash []byte
}

func (b *Block) Validate() error {
	hash := GenerateHash(b.Time.UnixNano(), b.Data, b.PrevHash)
	if !bytes.Equal(hash, b.Hash) {
		return fmt.Errorf("Hash is invalid \nit should be :\n%x \nbut is :\n%x", b.Hash, hash)
	}
	return nil
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"\n-------------\nTime : %s \nData : %s \nHash : %x \nPrevHash : %x\n",
		b.Time, b.Data, b.Hash, b.PrevHash)
}

func NewBlock(data string, prevHash []byte) *Block {
	b := Block{
		Data:     []byte(data),
		Time:     time.Now(),
		PrevHash: prevHash,
	}
	b.Hash = GenerateHash(b.Time.UnixNano(), b.Data, b.PrevHash)

	return &b

}
