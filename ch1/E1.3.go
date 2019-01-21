// TODO: add benchmark tests for different version of "echo"
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo()
	echo0()
}

func echo0() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
