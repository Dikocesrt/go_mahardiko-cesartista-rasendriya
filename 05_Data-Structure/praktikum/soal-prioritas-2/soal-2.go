package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(num)))
	for i := 3; i <= maxDivisor; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func groupPrimes(input []int) [][]int {
	grouped := make([][]int, 0)
	primes := make([]int, 0)
	nonPrimes := make([]int, 0)

	for _, num := range input {
		if isPrime(num) {
			primes = append(primes, num)
		} else {
			nonPrimes = append(nonPrimes, num)
		}
	}

	grouped = append(grouped, primes)
	grouped = append(grouped, nonPrimes)

	return grouped
}

func main() {
	input1 := []int{2, 3, 4, 5, 6, 7, 8}
	output1 := groupPrimes(input1)
	fmt.Println("Input:", input1)
	fmt.Println("Output:", output1)

	input2 := []int{15, 24, 3, 9, 5}
	output2 := groupPrimes(input2)
	fmt.Println("Input:", input2)
	fmt.Println("Output:", output2)

	input3 := []int{4, 8, 9, 12}
	output3 := groupPrimes(input3)
	fmt.Println("Input:", input3)
	fmt.Println("Output:", output3)
}
