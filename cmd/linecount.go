package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"
)

func init() {
	linecountCmd.Flags().StringP("file", "f", "", "File to operate on")
	linecountCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(linecountCmd)
}

func isValid(filename string) (bool, error) {

	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false, fmt.Errorf("error: No such file '%v'", filename)
	}

	if info.IsDir() {
		return false, fmt.Errorf("error: Expected file got directory '%v'", filename)
	}

	mime, err := mimetype.DetectFile(filename)

	if mime.Is("application/x-mach-binary") {
		return false, fmt.Errorf("error: Cannot do linecount for binary file '%v'", filename)
	}

	return true, err
}

func count(filename string) (int, error) {
	isValid, err := isValid(filename)

	if isValid {
		lines, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Errorf("error: Couldn't read file '%v'", filename)
		}
		return strings.Count(string(lines), "\n"), err
	}

	return 0, err
}

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Count lines in a file",
	Long:  `Linecount counts the number of lines in a text file. A valid text file is required.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		lines, err := count(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(lines)
		}
	},
}
