package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
        return 1
    } else if n%2 == 0 {
        halfPow := pow(x, n/2)
        return halfPow * halfPow
    } else {
        halfPow := pow(x, (n-1)/2)
        return halfPow * halfPow * x
    }

}

func main(){
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}