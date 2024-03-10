package main

import (
    "fmt"
    "math"
)

func isPrime(number int) bool {
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

func findPrimes(start, end int, primes chan int) {
    defer close(primes)
    for i := start; i <= end; i++ {
        if isPrime(i) {
            primes <- i
        }
    }
}

func powerPrimes(primes chan int, results chan int) {
    for prime := range primes {
        results <- int(math.Pow(float64(prime), 2))
    }
    close(results)
}

func main() {
    primes := make(chan int)
    results := make(chan int)

    go findPrimes(1, 100, primes)
    go powerPrimes(primes, results)

    for result := range results {
        fmt.Println(result)
    }
}
