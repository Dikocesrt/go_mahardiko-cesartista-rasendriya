package main

import (
	"fmt"
	"math"
)

func main() {
	primeFactors(20)   // 2, 2, 5
	primeFactors(75)   // 3, 5, 5
	primeFactors(1024) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {
	for n%2 == 0 {
		n = n / 2
		fmt.Print("2, ")
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		for n%i == 0 {
			n = n / i
			fmt.Printf("%d, ", i)
		}
	}

	if n > 2 {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}