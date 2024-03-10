package main

import (
	"fmt"
)

func PrintPerChar(input string, ch chan string) {
	for i := range input {
		ch <- input[:i+1]
	}
	close(ch)
}

func main() {
	input := "rusak"

	ch := make(chan string)

	go PrintPerChar(input, ch)

	for s := range ch {
		fmt.Println(s)
	}
}
