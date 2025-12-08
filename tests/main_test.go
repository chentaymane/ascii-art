package tests

import (
	"os"
	"strconv"
	"testing"

	"ascii-art/converter"
)

var tests = []string{ // https://github.com/01-edu/public/tree/master/subjects/ascii-art/audit
	"hello", "HELLO", "HeLlo HuMaN", "1Hello 2There", "Hello\nThere", "Hello\n\nThere", "{Hello & There #}",
	"hello There 1 to 2!", "MaD3IrA&LiSboN", `1a\"#FdwHywR&/()=`, "{|}~", `[\]^_ 'a`, "RGB", ":;<=>?@", `\!" #$%&'"'"'()*+,-./`,
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz",
}

func Test_main(t *testing.T) {
	for i, input := range tests {
		bytes, _ := os.ReadFile("cases/" + strconv.Itoa(i))
		output := string(bytes)
		font := "standard"
		result := converter.Run(input, "../banner/"+font)
		if result != output {
			t.Errorf("Test failed.\nWant:\n%s\nHas:\n%s", output, result)
		}
	}
}

func Test_supported(t *testing.T) {
	for i := 0; i < 256; i++ {
		font := "standard"
		result := converter.Run(string(rune(i)), "../banner/"+font)

		if i > '~' && result != "" {
			t.Error("Must return empty string.")
		} else if i > ' ' && i < '~' && result == "" {
			t.Error("Must not return empty string.")
		} else {
			// characters under ' ' are not printable
		}
	}
}
