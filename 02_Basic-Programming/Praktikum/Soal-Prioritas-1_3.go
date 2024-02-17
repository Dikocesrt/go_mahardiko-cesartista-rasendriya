package main

import "fmt"

func main(){
	for i := 1; i <= 100; i++ {
		if i % 4 == 0 {
			if i % 7 == 0 {
				fmt.Println("buzz")
			} else {
				fmt.Println(i * i)
			}
		} else if i % 7 == 0 {
			fmt.Println(i * i * i)
		} else {
			fmt.Println(i)
		}
	}
}