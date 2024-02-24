package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(number)))
	for i := 3; i <= maxDivisor; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(121))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}