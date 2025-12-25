package internal

import (
	"fmt"
	"os"
)

// PrintByteCount reads the file and returns byte count
func PrintByteCount(filename string) {
	// Open the file
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening file: %w", err)
		os.Exit(1)
	}

	fileSize := fileInfo.Size()
	fmt.Fprintf(os.Stdout, "%v %s", fileSize, filename)
	os.Exit(0)
}
