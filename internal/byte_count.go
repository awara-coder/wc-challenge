package internal

import (
	"fmt"
	"os"
)

// PrintByteCount reads the file and returns byte count
func PrintByteCount(filename string) {
	byteCount, err := getByteCount(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening file: %w", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%v ", byteCount)
}

func getByteCount(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, fmt.Errorf("error while opening file: %w", err)
	}

	fileSize := fileInfo.Size()
	return fileSize, nil
}
