package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

// struct untuk menyimpan hasil perhitungan frekuensi
type WordFreq struct {
	Word  string
	Count int
}

func main() {
	file, err := os.Open("document.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freqCh := make(chan WordFreq)
	doneCh := make(chan bool)

	var wg sync.WaitGroup

	wordFreqMap := make(map[string]int)

	countWord := func(text string) {
		words := strings.Fields(text)
		for _, word := range words {
			wordFreqMap[word]++
		}
		wg.Done()
	}

	sendFreq := func() {
		for word, count := range wordFreqMap {
			freqCh <- WordFreq{word, count}
		}
		close(freqCh)
	}

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			wg.Add(1)
			go countWord(line)
		}
		wg.Wait()
		sendFreq()
	}()

	go func() {
		outputFile, err := os.Create("output.csv")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer outputFile.Close()

		writer := csv.NewWriter(outputFile)
		defer writer.Flush()

		writer.Write([]string{"word", "frequencies"})

		for freq := range freqCh {
			writer.Write([]string{freq.Word, fmt.Sprintf("%d", freq.Count)})
		}
        fmt.Println("TEST B")
        doneCh <- true
	}()

	<-doneCh
	fmt.Println("Selesai menghitung frekuensi kata.")
}
