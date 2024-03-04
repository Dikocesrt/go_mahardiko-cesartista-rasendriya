package main

import "fmt"

func main() {
    fmt.Println(spinString("alta"))    // aalt
    fmt.Println(spinString("alterra")) // aalrtre
    fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
    chars := []rune(sentence)
    length := len(chars)
    
    for i := 0; i < length-1; i += 2 {
        temp := chars[i+1]
        chars[i+1] = chars[length-1]
        for j := length-1; j > i+2; j-- {
            if j < len(chars) {
                chars[j] = chars[j-1]
            }
        }
        if i+2 < len(chars) {
            chars[i+2] = temp
        }
    }
    
    return string(chars)
}
