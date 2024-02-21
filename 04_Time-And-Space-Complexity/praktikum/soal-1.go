package main

import "fmt"

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	} else if number == 2 {
		return true
	} else {
		if number % 2 == 0 {
			return false
		} else {
			if number % 3 == 0 {
				return false
			} else if number % 5 == 0 {
				return false
			} else if number % 7 == 0 {
				return false
			}
		}
	}
	return true
}

func main(){
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}