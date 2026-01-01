/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/awara-coder/wc-challenge/internal"
	"github.com/spf13/cobra"
)

var (
	printByteCount bool
	printLineCount bool
	printWordCount bool
)

const (
	printByteCountFlagName = "byte"
	printLineCountFlagName = "line"
	printWordCountFlagName = "word"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc [-c | -l | -w] file",
	Short: "word, line, character and byte count", // Same as original wc
	Long: `This is an imitation of wc command line tool  written in golang.
Used to find word, line, character and byte count of file.
Currently supporting only single file.`,
	Args: internal.ValidateFileArgs,

	Run: func(cmd *cobra.Command, args []string) {
		// Call the right command based on flag
		switch {
		case printByteCount:
			internal.PrintByteCount(args[0])
		case printLineCount:
			internal.PrintLineCount(args[0])
		case printWordCount:
			internal.PrintWordCount(args[0])

		}
		// Print file name
		fmt.Printf("%v", args[0])

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Byte count flag: -c or --byte
	rootCmd.PersistentFlags().BoolVarP(&printByteCount, printByteCountFlagName, "c", false, "Returns number of bytes")
	rootCmd.PersistentFlags().BoolVarP(&printLineCount, printLineCountFlagName, "l", false, "Returns number of lines")
	rootCmd.PersistentFlags().BoolVarP(&printWordCount, printWordCountFlagName, "w", false, "Returns number of words")

	// Make sure both of them are mutually exclusive
	rootCmd.MarkFlagsMutuallyExclusive(printByteCountFlagName, printLineCountFlagName, printWordCountFlagName)

	// Make sure atleast one command is set
	rootCmd.MarkFlagsOneRequired(printByteCountFlagName, printLineCountFlagName, printWordCountFlagName)
}
