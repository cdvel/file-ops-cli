package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fops " + Version)
	},
}
