package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func getInfo(student []Student) {
	sort.Slice(student, func (i, j int) bool {
		return student[i].MathScore > student[j].MathScore
	})
	fmt.Printf("best student in math : %s with %d\n", student[0].Name, student[0].MathScore)

	sort.Slice(student, func (i, j int) bool {
		return student[i].ScienceScore > student[j].ScienceScore
	})
	fmt.Printf("best student in math : %s with %d\n", student[0].Name, student[0].ScienceScore)

	sort.Slice(student, func (i, j int) bool {
		return student[i].EnglishScore > student[j].EnglishScore
	})
	fmt.Printf("best student in math : %s with %d\n", student[0].Name, student[0].EnglishScore)

	var average, max int
	var nama string
	for i := 0; i < len(student); i++ {
		average = (student[i].MathScore + student[i].ScienceScore + student[i].EnglishScore) / 3
		if i == 0 {
			max = average
			nama = student[i].Name
		} else {
			if average > max {
				max = average
				nama = student[i].Name
			}
		}
	}
	fmt.Printf("best student in average : %s with %d", nama, max)
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
}