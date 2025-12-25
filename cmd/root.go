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
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc [-c] file",
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
		default:
			fmt.Fprintf(os.Stderr, "no flag set!")
		}
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
	rootCmd.PersistentFlags().BoolVarP(&printByteCount, "byte", "c", false, "Returns number of bytes")
	rootCmd.MarkPersistentFlagRequired("byte")
}
