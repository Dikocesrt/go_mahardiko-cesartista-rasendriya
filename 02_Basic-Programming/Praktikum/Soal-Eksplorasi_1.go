package main

import "fmt"

func main(){
	counterPertama := 0
	counterKedua := 0

	isAnagram := true

	pertama := "kasur"
	kedua := "rusak"

	if len(pertama) == len(kedua) {
		for i := 0; i < len(pertama); i++ {
			temp := string(pertama[i])
			for j := i+1; j < len(pertama); j++ {
				if temp == string(pertama[j]) {
					counterPertama += 1
					pertama = pertama[:j] + pertama[j+1:]
				}
			}
			for k := 0; k < len(kedua); k++ {
				if temp == string(kedua[k]) {
					counterKedua += 1
				}
			}
			counterPertama += 1
			if counterPertama != counterKedua {
				isAnagram = false
			}
			counterPertama = 0
			counterKedua = 0
		}
	} else {
		isAnagram = false
	}

	if(isAnagram){
		fmt.Println("Anagram")
	} else {
		fmt.Println("Bukan Anagram")
	}
}