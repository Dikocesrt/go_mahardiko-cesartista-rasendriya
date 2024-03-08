package main

import "fmt"

var memo map[int]int

func fibonacci(number int) int {
    memo = make(map[int]int)
    return fib(number)
}

func fib(n int) int {
    if n <= 1 {
        return n
    }
    if _, ok := memo[n]; !ok {
        memo[n] = fib(n-1) + fib(n-2)
    }
    return memo[n]
}

func main() {
    fmt.Println(fibonacci(0))  // 0
    fmt.Println(fibonacci(2))  // 1
    fmt.Println(fibonacci(9))  // 34
    fmt.Println(fibonacci(10)) // 55
    fmt.Println(fibonacci(12)) // 144
}
