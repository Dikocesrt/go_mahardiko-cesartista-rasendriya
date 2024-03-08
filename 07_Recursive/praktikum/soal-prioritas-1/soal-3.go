package main

import "fmt"

func main() {
	fmt.Println(getSequence(4))   // 10
	fmt.Println(getSequence(15))  // 120
	fmt.Println(getSequence(100)) // 5050
}

func getSequence(n int) int {
	var deret []int
	beda := 1
	for i := 0; i <= n; i++ {
		if i == 0 {
			deret = append(deret, 0)
		} else {
			deret = append(deret, deret[i-1] + beda)
		}
		beda = i+1
	}
	return deret[n]
}