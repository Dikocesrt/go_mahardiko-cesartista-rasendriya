package main

import (
    "fmt"
)

type Stack struct {
    items map[string]bool
}

func NewStack() *Stack {
    return &Stack{items: make(map[string]bool)}
}

func (s *Stack) push(item string) {
    if _, exists := s.items[item]; !exists {
        s.items[item] = true
    }
}

func (s *Stack) pop() string {
    var poppedItem string
    for item := range s.items {
        poppedItem = item
        delete(s.items, item)
        break
    }
    return poppedItem
}

func (s *Stack) isEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack) values() []string {
    var stackValues []string
    for item := range s.items {
        stackValues = append(stackValues, item)
    }
    return stackValues
}

func main() {
    stack := NewStack()

    // Menambahkan data ke dalam stack
    stack.push("a")
    stack.push("b")
    stack.push("c")
    stack.push("a") // Ini tidak akan ditambahkan karena data sudah ada

    // Mengambil dan menampilkan data dari stack
    fmt.Println("Data yang diambil dari stack:")
    for !stack.isEmpty() {
        fmt.Println(stack.pop())
    }
}
