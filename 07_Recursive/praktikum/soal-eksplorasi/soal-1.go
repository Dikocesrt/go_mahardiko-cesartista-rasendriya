package main

import (
	"fmt"
)

func main() {
	fmt.Println(getDiamondSeq(4))   // 16
	fmt.Println(getDiamondSeq(25))  // 625
	fmt.Println(getDiamondSeq(100)) // 10000
}

func getDiamondSeq(n int) int {
	var deret []int
	beda := 1
	for i := 0; i <= n; i++ {
		if i == 0 {
			deret = append(deret, 0)
		} else {
			deret = append(deret, deret[i-1] + beda)
			beda += 2
		}
	}
	return deret[n]
}