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
	bc := NewBlockchain(7)
	//bc.Add("Second Block on Blockchain")
	//bc.Add("Third Block on Blockchain")

	//bc.Blocks[1].Time = time.Now() // for check invalid hash uncomment this line

	if err := bc.Validate(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(bc)
}
