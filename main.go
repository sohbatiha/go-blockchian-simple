package main

import "fmt"

func main() {
/*	b:= NewBlock("Hi Blockchain")

	//b.Time = time.Now() // for check invalid hash uncomment this line
	if err := b.Validate(); err!= nil {
		log.Fatal(err.Error())
	}*/

	bc:=NewBlockchain()
	bc.Add("Second Block on Blockchain")
	bc.Add("Third Block on Blockchain")
	fmt.Println(bc)
}
