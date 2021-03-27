package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fops",
	Short: "A simple yet powerful file checker CLI",
	Long: `fops is a simple CLI for file operations.
It supports line count and checksum operations on valid non binary files.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
