package main

import (
	"fmt"
	"slices"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("sum: %.2f\n", sum(data1))       // 71.00
	fmt.Printf("mean: %.2f\n", mean(data1))     // 5.46
	fmt.Printf("median: %.2f\n", median(data1)) // 5.00
	fmt.Printf("mode: %.2f\n", mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("sum: %.2f\n", sum(data2))       // 103.00
	fmt.Printf("mean: %.2f\n", mean(data2))     // 7.92
	fmt.Printf("median: %.2f\n", median(data2)) // 8.00
	fmt.Printf("mode: %.2f\n", mode(data2))     // 1.00
}

func sum(data []float64) float64 {
	result := 0.0
	for _, value := range data {
		result += value
	}
	return result
}

func mean(data []float64) float64 {
	size := len(data)
	result := sum(data)/float64(size)
	return result
}

func median(data []float64) float64 {
	var idx int
	if len(data) % 2 == 0 {
		idx = ((len(data) / 2) + ((len(data) / 2) + 1)) / 2
	} else {
		idx = (len(data) + 1) / 2
	}
	slices.Sort[[]float64](data)
	return data[idx]
}

func mode(data []float64) float64 {
	encountered := map[float64]bool{}

	var modusTemp float64
	var counter int
	var counterModus int

	for i, value := range data {
		counter = 0
		if !encountered[value] {
			encountered[value] = true
			for j := i+1; j < len(data); j++ {
				if value == data[j] {
					counter++
				}
			}
			if i == 0 {
				modusTemp = value
				counterModus = counter
			} else {
				if counterModus < counter {
					modusTemp = value
					counterModus = counter
				}
			}
		}
	}
	return modusTemp
}