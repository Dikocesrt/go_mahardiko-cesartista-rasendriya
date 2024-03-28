package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Words struct {
	Word string `json:"word"`
	Length int `json:"length"`
	NumOfVocals int `json:"num_of_vocals"`
	Palindrom bool `json:"palindrom"`
}

type BaseResponseError struct {
	Message   string `json:"message"`
}

type BaseResponseSuccess struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.POST("/words", addwordHandler)
	e.GET("/words", getAllWords)
	e.Start(":8080")
}

func getAllWords(c echo.Context) error {
	words := generateWords()

	baseResponseSuccess := BaseResponseSuccess {
		Data: words,
		Message: "all words",
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func addwordHandler(c echo.Context) error {
	var word Words
	c.Bind(&word)

	errResponse := BaseResponseError{}
	if word.Word == "" {
		errResponse.Message = "failed to add word data"
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	word.Length = len(word.Word)
	word.NumOfVocals = getNumberOfVocals(word.Word)
	word.Palindrom = isPolindrom(word.Word)

	baseResponseSuccess := BaseResponseSuccess {
		Message: "Word Added",
		Data: word,
	}

	return c.JSON(http.StatusCreated, baseResponseSuccess)
}

func getNumberOfVocals(word string) int {
	word = strings.ToLower(word)
	counter := 0;
	vocals := "aiueo"
	for _, char := range word {
		for _, vocal := range vocals {
			if char == vocal {
				counter++
			}
		}
		
	}
	return counter
}

func reverseWord(word string) string {
	wordBytes := []byte(word)

	length := len(wordBytes)

	for i := 0; i < length/2; i++ {
		wordBytes[i], wordBytes[length-i-1] = wordBytes[length-i-1], wordBytes[i]
	}

	return string(wordBytes)
}

func isPolindrom(word string) bool {
	word = strings.ToLower(word)
	reverseWord := reverseWord(word)

	if word == reverseWord {
		return true
	}

	return false
}

func generateWords() []Words {
	words := []Words{
		{"civic", len("civic"), getNumberOfVocals("civic"), isPolindrom("civic")},
		{"kAsUr RuSaK", len("kAsUr RuSaK"), getNumberOfVocals("kAsUr RuSaK"), isPolindrom("kAsUr RuSaK")},
		{"rumah", len("rumah"), getNumberOfVocals("rumah"), isPolindrom("rumah")},
	}

	return words
}