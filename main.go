package main

import (
	"fmt"
	"os"

	"main/converter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: invalid number of inputs")
		return
	}
	if len(os.Args[1]) == 0 {
		return
	}

	fmt.Print(converter.Run(os.Args[1], "standard.txt"))
}
