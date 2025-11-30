package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: invalid number of inputs")
		return
	}

	input := os.Args[1]

	linesInput := strings.Split(input, "\\n")

	font := "standard.txt"
	content, err := os.ReadFile(font)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	fontLines := strings.Split(text, "\n")

	for _, line := range linesInput {

		if line == "" {
			fmt.Println()
			continue
		}
		runes := []rune(line)
		chars := make([][]string, len(runes))

		for i, char := range runes {
			if char < ' ' || char > '~' {
				fmt.Println("unsupported character")
				return
			}

			index := int(((char - ' ') * 9) + 1)
			chars[i] = fontLines[index : index+8]
		}

		for height := 0; height < 8; height++ {
			for i := range chars {
				fmt.Print(chars[i][height])
			}
			fmt.Println()
		}

	}
}
