package converter

import (
	"fmt"
	"os"
	"strings"
)

func Run(input string, font_path string) string {
	linesInput := strings.Split(input, "\\n")
	final := ""

	font := "standard.txt"
	content, err := os.ReadFile(font)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// fmt.Printf("%q", content[:12])
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	// fmt.Printf("%q", text[:12])

	fontLines := strings.Split(text, "\n")
	// fmt.Println(len(linesInput))
	for _, line := range linesInput {

		if line == "" {
			final += "\n"
			continue
		}
		runes := []rune(line)
		chars := make([][]string, len(runes))

		for i, char := range runes {
			if char < ' ' || char > '~' {
				fmt.Println("unsupported character")
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
