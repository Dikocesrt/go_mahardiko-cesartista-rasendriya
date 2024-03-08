package main

import (
	"fmt"
	"strconv"
)

func toBinary(bilangan int) []int {
	var hasil []int
	for i := 0; i <= bilangan; i++ {
		temp := strconv.FormatInt(int64(i), 2)
		num, _ := strconv.Atoi(temp)
		hasil = append(hasil, num)
	}
	return hasil
}

func main() {
	fmt.Println(toBinary(2))
	fmt.Println(toBinary(3))
}