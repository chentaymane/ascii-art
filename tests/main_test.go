package tests

import (
	"os"
	"strings"
	"testing"

	"main/converter"
)

func Test_main(t *testing.T) {
	files, err := os.ReadDir("cases")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() == false {
			bytes, _ := os.ReadFile("cases/" + file.Name())
			output := string(bytes)
			input := strings.Split(file.Name(), ".")[0]
			result := converter.Run(input, "../standard.txt")
			if result != output {
				t.Errorf("Test failed.\nWant:\n%s\nHas:\n%s", output, result)
			}
		}
	}
}
