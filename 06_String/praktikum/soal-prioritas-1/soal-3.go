package main

import (
    "fmt"
)

func main() {
    fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
    fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
    fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) [][]string {
    var palindromes [][]string
    var nonPalindromes []string

    for _, word := range words {
        if isPalindrome(word) {
            palindromes = append(palindromes, []string{word})
        } else {
            nonPalindromes = append(nonPalindromes, word)
        }
    }

    groupedPalindromes := groupPalindromes(palindromes)

    return append(groupedPalindromes, nonPalindromes)
}

func isPalindrome(word string) bool {
    for i := 0; i < len(word)/2; i++ {
        if word[i] != word[len(word)-1-i] {
            return false
        }
    }
    return true
}

func groupPalindromes(words [][]string) [][]string {
    grouped := make(map[string][]string)

    for _, wordGroup := range words {
        sortedWord := sortString(wordGroup[0])
        grouped[sortedWord] = append(grouped[sortedWord], wordGroup[0])
    }

    var result [][]string
    for _, group := range grouped {
        result = append(result, group)
    }

    return result
}

func sortString(s string) string {
    runes := []rune(s)
    for i := 0; i < len(runes)-1; i++ {
        for j := i + 1; j < len(runes); j++ {
            if runes[i] > runes[j] {
                runes[i], runes[j] = runes[j], runes[i]
            }
        }
    }
    return string(runes)
}
