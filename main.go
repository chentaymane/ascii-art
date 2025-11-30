package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("eror number of inputs non valide ")
		return
	}
//	input:= os.Args[1]
font := "standard.txt"
content ,err := os.ReadFile(font)
if err != nil {
	fmt.Println(err)
	return
}
lines := strings.Split(string(content),"\n")
		fmt.Println(lines)

	//fmt.Println(string('~'-('~'-'!')))


}