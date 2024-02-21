package main

import "fmt"

func main(){
	bilangan := 26
	
	for i := 1; i <= bilangan; i++ {
		if bilangan % i == 0 {
			if i % 2 == 0 {
				fmt.Printf("I")
			} else {
				fmt.Printf("O")
			}
		}
	}
}