package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println(pickVocals("alterra academy"))     // aea aae
    fmt.Println(pickVocals("sepulsa mantap jiwa")) // eua aa ia
    fmt.Println(pickVocals("kopi susu"))           // oi uu
}

func pickVocals(sentence string) string {
    vowels := "aeiou"
    var result strings.Builder

    for _,char := range sentence {
        if char == ' ' {
            result.WriteRune(' ')
        }
        
        if strings.ContainsRune(vowels, char) {
            result.WriteRune(char)
        }
    }

    return result.String()
}
