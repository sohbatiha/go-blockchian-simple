package main

type Blockchain struct {
	Blocks []*Block
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

	bc.Blocks = append(bc.Blocks, NewBlock(data, prevHash))

}


func NewBlockchain() *Blockchain {
	bc := Blockchain{
		Blocks: []*Block{NewBlock("Genesis Block", []byte{})},
	}

	return &bc
}

