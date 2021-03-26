package cmd

import (
	"fmt"
	. "fops-cli/algorithm"
	. "fops-cli/filesys"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var hashFunctions map[string]Algorithm

func init() {

	checksumCmd.Flags().StringP("file", "f", "", "File to operate on")

	checksumCmd.Flags().Bool("md5", false, "checksum using md5")
	checksumCmd.Flags().Bool("sha1", false, "checksum using sha1")
	checksumCmd.Flags().Bool("sha256", false, "checksum using sha256")

	checksumCmd.MarkFlagRequired("file")

	hashFunctions = map[string]Algorithm{
		"md5":    GetMD5(),
		"sha1":   GetSHA1(),
		"sha256": GetSHA256(),
	}

	rootCmd.AddCommand(checksumCmd)
}

func GetChecksum(value []byte, hashFn string) string {
	return hashFunctions[hashFn].GetChecksum(value)
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Generate checksum of a file",
	Long: `Checksum retrieves the cryptographic hash of a file.
A valid text file and an algorithm flag (MD5, SHA1, SHA256) are needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")

		fs := File(filename).CheckFile()

		if fs.Err != nil {
			fmt.Println(fs.Err)
			return
		}

		contents, _ := ioutil.ReadFile(filename)
		algo := ""

		for _, algorithm := range hashFunctions {
			match, _ := cmd.Flags().GetBool(algorithm.GetType())
			if match {
				algo = algorithm.GetType()
			}
		}

		fmt.Println(GetChecksum(contents, algo))

	},
}
