package main

import "fmt"

func main() {
    fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    // [1,5,2,4,3]
    fmt.Println(spinSlice([]int{6, 7, 8}))          // [6,8,7]
    fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) // [8,1,5,2,4,3]
}

func spinSlice(numbers []int) []int {
    newSlice := []int{}
    bantu := len(numbers) - 1
    bantu2 := 1
    for i := 0; i < len(numbers); i++ {
        if i == 0 {
            newSlice = append(newSlice, numbers[i])
        } else if i % 2 == 0 {
            newSlice = append(newSlice, numbers[i-bantu2])
            bantu2+=1
        } else {
            newSlice = append(newSlice, numbers[bantu])
            bantu-=1
        }
    }
    return newSlice
}
