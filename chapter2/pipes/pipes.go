package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount takes a file and returns a map
// with each word as a key and it's number of
// appearances as a value
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	// make a scanner to work on the file
	// io.Reader interface
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
