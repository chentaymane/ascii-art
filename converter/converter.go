package converter

import (
	"fmt"
	"os"
	"strings"
)

func Run(input string, font string) string {
	if len(input) == 0 {
		return ""
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	linesInput := strings.Split(input, "\n")
	if linesInput[0] == "" && linesInput[1] == "" {
		linesInput = linesInput[1:] // remove extra element from strings.Split
	}

	content, err := os.ReadFile(font + ".txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	text := strings.ReplaceAll(string(content), "\r\n", "\n")

	fontLines := strings.Split(text, "\n")
	final := ""
	for _, line := range linesInput {
		if line == "" {
			final += "\n"
			continue
		}
		runes := []rune(line)
		chars := make([][]string, len(runes))

		for i, char := range runes {
			if char < ' ' || char > '~' {
				fmt.Println("Error: unsupported character.")
				return ""
			}

			index := int(((char - ' ') * 9) + 1)
			chars[i] = fontLines[index : index+8]
		}

		for height := 0; height < 8; height++ {
			for i := range chars {
				final += chars[i][height]
			}
			final += "\n"
		}
	}
	return final
}
