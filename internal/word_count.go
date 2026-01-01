package internal

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var whiteSpaceCharacters = []byte{' ', '\t', '\n', '\v', '\f', '\r'}

func PrintWordCount(filename string) {
	byteCount, err := getWordCount(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening file: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%v ", byteCount)
}

func getWordCount(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error while opening file: %w", err)
	}

	wordCount := int64(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		currentLineWordCount := int64(0)

		bytes := scanner.Bytes()
		inWord := 0
		for _, byte := range bytes {
			// is whitespace character
			if slices.Contains(whiteSpaceCharacters, byte) {
				currentLineWordCount += int64(inWord)
				inWord = 0
			} else {
				inWord = 1
			}
		}
		// Handle last word
		currentLineWordCount += int64(inWord)
		fmt.Println(currentLineWordCount)
		wordCount += int64(currentLineWordCount)
	}

	if scanner.Err() != nil {
		return 0, fmt.Errorf("error while scannling files for word count: %w", err)
	}

	return wordCount, nil
}
