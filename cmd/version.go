package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

func SetVersion(newVersion string) {
	Version = newVersion
}

func GetVersion() string {
	return Version
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fops " + Version)
	},
}
