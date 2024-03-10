package main

import (
	"fmt"
	"math"
)

func isComposite(n int) bool {
    if n < 4 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return true
        }
    }
    return false
}

func findComposites(start, end int, composites chan int) {
    defer close(composites)
    for i := start; i <= end; i++ {
        if isComposite(i) {
            composites <- i
        }
    }
}

func powerAndCheck(composites chan int, results chan string) {
    for composite := range composites {
        power := int(math.Pow(float64(composite), 2))
        if power%2 == 0 {
            results <- "Ping"
        } else {
            results <- "Pong"
        }
    }
    close(results)
}

func main() {
    composites := make(chan int)
    results := make(chan string)

    go findComposites(1, 100, composites)
    go powerAndCheck(composites, results)

    for result := range results {
        fmt.Println(result)
    }
}
