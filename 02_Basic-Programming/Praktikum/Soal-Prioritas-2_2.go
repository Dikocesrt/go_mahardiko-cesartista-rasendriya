package main

import "fmt"

func main(){
	budget := 51
	waktuPengerjaan := 10
	tingkatKesulitan := 3
	skor := 0

	if budget >= 0 && budget <= 50 {
		skor += 4
	} else if budget >= 51 && budget <= 80 {
		skor += 3
	} else if budget >= 81 && budget <= 100 {
		skor += 2
	} else if budget > 100 {
		skor += 1
	}

	if waktuPengerjaan >= 0 && waktuPengerjaan <= 20 {
		skor += 10
	} else if waktuPengerjaan >= 21 && waktuPengerjaan <= 30 {
		skor += 5
	} else if waktuPengerjaan >= 31 && waktuPengerjaan <= 50 {
		skor += 2
	} else if waktuPengerjaan > 50 {
		skor += 1
	}

	if tingkatKesulitan >= 0 && tingkatKesulitan <= 3 {
		skor += 10
	} else if tingkatKesulitan >= 4 && tingkatKesulitan <= 6 {
		skor += 5
	} else if tingkatKesulitan >= 8 && tingkatKesulitan <= 10 {
		skor += 1
	} else if tingkatKesulitan > 10 {
		skor += 0
	}

	if skor >= 17 && skor <= 24 {
		fmt.Println("High")
	} else if skor >= 10 && skor <= 16 {
		fmt.Println("Medium")
	} else if skor >= 3 && skor <= 9 {
		fmt.Println("Low")
	} else if skor <= 2 {
		fmt.Println("Impossible")
	}
}