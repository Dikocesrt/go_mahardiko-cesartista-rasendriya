package main

import (
	"fmt"
)

func fibX(n int) int {
	var fibonacci []int
	var result int

	for i := 0; i <= n; i++ {
		if i == 0 {
			fibonacci = append(fibonacci, 0)
		} else if i == 1 {
			fibonacci = append(fibonacci, 1)
		} else {
			fibonacci = append(fibonacci, fibonacci[i-2] + fibonacci[i-1])
		}
		result += fibonacci[i]
	}
	return result
}

func main() {
	fmt.Println(fibX(5))
}