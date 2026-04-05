package main

import (
	"ascii-art/converter"
	"fmt"
	"os"
)

func main() {
	text, font, file, ok := converter.ParseArgs(os.Args)
	if !ok {
		fmt.Println(converter.ErrMsg)
		return
	}

	err := converter.Run(text, "./banner/"+font, file)
	if err != "" {
		fmt.Println(converter.ErrMsg)
	}
}
