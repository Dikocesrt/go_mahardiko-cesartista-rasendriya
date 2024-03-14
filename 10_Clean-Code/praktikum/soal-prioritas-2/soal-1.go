package main

import "fmt"

type set struct {
	data map[int]int
}

func (setObject set) add(key int) {
	setObject.data[key] = 1
}

func (setObject set) get() []int {
	result := []int{}

	for key := range setObject.data {
		result = append(result, key)
	}

	return result
}

func (setObject set) remove(key int) {
	delete(setObject.data, key)
}

func main() {
	setObject := set{
		data: map[int]int{},
	}

	setObject.add(1)
	setObject.add(2)
	setObject.add(1)
	setObject.add(3)
	setObject.add(4)
	setObject.add(5)

	setObject.remove(100)

	fmt.Println(setObject.get())
}