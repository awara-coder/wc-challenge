package internal

import (
	"bufio"
	"fmt"
	"os"
)

func PrintLineCount(filename string) {
	byteCount, err := getLineCount(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening file: %w", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%v %s", byteCount, filename)
	os.Exit(0)
}

func getLineCount(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error while opening file: %w", err)
	}

	lineCount := int64(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if scanner.Err() != nil {
		return 0, fmt.Errorf("error while scannling files for line count: %w", err)
	}

	return lineCount, nil
}
