package cmd

import (
	"fmt"
	. "fops-cli/filesys"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

var linecount *cobra.Command

func init() {
	LinecountCmd().Flags().StringP("file", "f", "", "File to operate on")
	LinecountCmd().MarkFlagRequired("file")
	rootCmd.AddCommand(LinecountCmd())
}

func count(filename string) (int, error) {

	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("error: Couldn't read file '%v'", filename)
	}
	return strings.Count(string(lines), "\n"), err
}

func LinecountRun(cmd *cobra.Command, args []string) error {
	filename, _ := cmd.Flags().GetString("file")
	fs := File(filename).CheckFile().CheckNotBinary()

	if fs.Err != nil {
		fmt.Println(fs.Err)
	} else {
		lines, err := count(fs.Filename)
		if err != nil {
			fmt.Println(err)
			return err
		} else {
			fmt.Println(lines)
		}
	}
	return nil
}

func LinecountCmd() *cobra.Command {
	if linecount == nil {
		linecount = &cobra.Command{
			Use:   "linecount",
			Short: "Count lines in a file",
			Long:  `Linecount counts the number of lines in a text file. A valid text file is required.`,
			RunE:  LinecountRun,
		}
	}
	return linecount
}
