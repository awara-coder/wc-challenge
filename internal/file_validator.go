package internal

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func ValidateFileArgs(cmd *cobra.Command, args []string) error {
	// We are supporting only 1 file for now.
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	// Make sure that the file exists for the given path.
	if !fileExists(args[0]) {
		return errors.New("file doesn't exist")
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	// check if no error and that file is not a directory
	return err == nil && !info.IsDir()
}
