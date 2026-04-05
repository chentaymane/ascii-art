package converter

import (
	"fmt"
	"os"
	"strings"
)

const 	ErrMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"


func Run(input string, font string, file string) string {
	if len(input) == 0 {
		return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	linesInput := strings.Split(input, "\n")
	if isOnlyNewline(input) {
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
				return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
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
	if file != "" {
		if strings.HasPrefix(file, "--output=") && strings.HasSuffix(file, ".txt") {
			file = file[len("--output="):]
			if file != "" {

				data := []byte(final)
				// 0644 sets the file permissions (read and write for owner, read for others)
				err := os.WriteFile(file, data, 0o644)
				if err != nil {
					fmt.Print(err)
				}
			} else {
				return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
			}
		} else {
			return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
		}
	} else {
		fmt.Print(final)
	}
	return ""
}

