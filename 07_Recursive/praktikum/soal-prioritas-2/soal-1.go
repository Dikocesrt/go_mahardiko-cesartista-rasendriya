package main

import (
	"fmt"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) int {
	return faktorial(2 * n) / (faktorial(n + 1) * faktorial(n))
}

func faktorial(a int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return a * faktorial(a-1)
	}
}