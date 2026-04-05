package converter

import (
	"os"
	"strings"
)

func ParseArgs(args []string) (text, font, file string, ok bool) {

	switch len(os.Args) {

	case 2:

		text = os.Args[1]
	case 3:
		if strings.HasPrefix(os.Args[1], "--") {
			text = args[2]
			file = os.Args[1]
		} else {
			text = args[1]
			font = os.Args[2]
		}
	case 4:
		file = os.Args[1]
		text = os.Args[2]
		font = os.Args[3]
	default:
		return "", "", "", false
	}

	font = checkFont(font)
	return text, font, file, true
}

func checkFont(font string) string {
	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		return "standard"
	}
	return font
}

func isOnlyNewline(s string) bool {
	for _, char := range s {
		if char != '\n' {
			return false
		}
	}
	return true
}
