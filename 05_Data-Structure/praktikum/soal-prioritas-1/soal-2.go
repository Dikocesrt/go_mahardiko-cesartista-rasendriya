package main

import "fmt"

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
	result := 0
	for _, value := range numbers {
		if value > 1 {
			prime := true

			for i := 2; i*i <= value; i++ {
				if value % i == 0 {
					prime = false
					break
				}
			}

			if prime {
				result += value
			}
		}
	}
	return result
}