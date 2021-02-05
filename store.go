package main

type Store interface {
	Load(hash []byte) (*Block, error)
	LastHash() ([]byte, error)
	Append(block *Block) error
}

func Iterate(store Store, fn func(block *Block) error) error {

	lastHash, err := store.LastHash()

	if err != nil {
		return err
	}

	for {
		block, err := store.Load(lastHash)
		if err != nil {
			return err
		}
		if err := fn(block); err != nil {
			return err
		}

		if len(block.PrevHash) <= 0 {
			return nil
		}

		lastHash = block.PrevHash
	}

}
