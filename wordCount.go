package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type WordCount struct {
	Word  string
	Count int
}

func wordCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		searchAllWords(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func searchAllWords(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./uploads")
	if err != nil {
		return
	}

	wordCounts := make(map[string]int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			path := filepath.Join(filepath.Join("./uploads", filename))

			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Failed to open file %s: %s\n", path, err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				word := strings.ToLower(scanner.Text())
				mu.Lock()
				wordCounts[word]++
				mu.Unlock()
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error scanning file %s: %s\n", path, err)
			}
		}(file.Name())
	}

	wg.Wait()

	topWords := make([]WordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		topWords = append(topWords, WordCount{Word: word, Count: count})
	}

	totalWordCount := 0
	for _, wordCount := range topWords {
		totalWordCount += wordCount.Count
	}
	message := fmt.Sprintf("Total word count: %d\n", totalWordCount)
	fmt.Fprint(w, message)
}
