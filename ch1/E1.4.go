// Modify dup2 to print the names of all files in which each duplicated
// line occurs.
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := []string{}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts) {
				names = append(names, f.Name())
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println(names)
}

func countLines(f *os.File, counts map[string]int) bool {
	dupInFile := false
	input := bufio.NewScanner(f)
	temp := make(map[string]int)
	for input.Scan() {
		temp[input.Text()]++
	}
	for k, v := range temp {
		if v > 1 {
			dupInFile = true
		}
		counts[k] += v
	}
	return dupInFile
	// NOTE: ignoring potential errors from input.Err()
}

//!-
