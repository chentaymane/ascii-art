package main

import (
	"fmt"
	"os"

	"ascii-art/converter"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	font := "standard"
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "thinkertoy", "shadow":
			font = os.Args[2]
		}
	}
	fmt.Print(converter.Run(os.Args[1], "./banner/"+font))
}
