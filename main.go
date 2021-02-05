package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	// difficulty is number of zeros
	difficulty := 4

	// store in memory
	//store := createNewMemoryStore()

	//store in file
	store := CreateNewFileStore("/home/mahdi/bc")

	bc, err := NewBlockchain(difficulty, store)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = bc.Add("Second Block on Blockchain")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = bc.Add("Third Block on Blockchain")
	if err != nil {
		log.Fatal(err.Error())
	}
	//bc.Blocks[1].Time = time.Now() // for check invalid hash uncomment this line

	if err := bc.Validate(); err != nil {
		log.Fatal(err.Error())
	}

	bc.Print()
}
