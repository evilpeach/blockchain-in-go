package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start!!")

	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
