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
    // Membuka file teks
    file, err := os.Open("document.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Membaca isi file baris per baris
    scanner := bufio.NewScanner(file)

    // Channel untuk mengirim hasil perhitungan frekuensi
    freqCh := make(chan WordFreq)
    // Channel untuk memberi tahu bahwa semua goroutine telah selesai
    doneCh := make(chan bool)

    // WaitGroup untuk menunggu selesai semua goroutine
    var wg sync.WaitGroup

    // Map untuk menyimpan frekuensi kata
    wordFreqMap := make(map[string]int)

    // Function untuk menghitung frekuensi kata
    countWord := func(text string) {
        defer wg.Done()
        words := strings.Fields(text)
        for _, word := range words {
            wordFreqMap[word]++
        }
    }

    // Function untuk mengirim hasil perhitungan ke dalam channel
    sendFreq := func() {
        defer close(freqCh)
        for word, count := range wordFreqMap {
            freqCh <- WordFreq{word, count}
        }
    }

    // Goroutine untuk membaca dan menghitung frekuensi kata dari setiap baris teks
    go func() {
        for scanner.Scan() {
            line := scanner.Text()
            wg.Add(1)
            go countWord(line)
        }
        wg.Wait()
        // Setelah selesai membaca file, kirim hasil perhitungan ke dalam channel
        go sendFreq()
    }()

    // Goroutine untuk menulis hasil perhitungan ke dalam file CSV
    go func() {
        defer close(doneCh)
        outputFile, err := os.Create("output.csv")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        defer outputFile.Close()

        writer := csv.NewWriter(outputFile)
        defer writer.Flush()

        // Tulis header ke dalam file CSV
        writer.Write([]string{"word", "frequencies"})

        // Menerima hasil perhitungan dari channel dan menuliskannya ke dalam file CSV
        for freq := range freqCh {
            writer.Write([]string{freq.Word, fmt.Sprintf("%d", freq.Count)})
        }
    }()

    // Menunggu selesai semua goroutine
    <-doneCh
    fmt.Println("Selesai menghitung frekuensi kata.")
}
