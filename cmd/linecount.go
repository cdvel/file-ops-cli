package cmd

import (
	"fmt"
	. "fops-cli/filesys"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func init() {
	linecountCmd.Flags().StringP("file", "f", "", "File to operate on")
	linecountCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(linecountCmd)
}

func count(filename string) (int, error) {

	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("error: Couldn't read file '%v'", filename)
	}
	return strings.Count(string(lines), "\n"), err
}

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Count lines in a file",
	Long:  `Linecount counts the number of lines in a text file. A valid text file is required.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		fs := FileStatus{false, nil, filename}.Validate()

		if fs.Err != nil {
			fmt.Println(fs.Err)
		} else {
			lines, err := count(fs.Filename)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(lines)
			}
		}
	},
}
