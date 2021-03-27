package cmd

import (
	"fmt"
	. "fops-cli/algorithm"
	. "fops-cli/filesys"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var hashFunctions map[string]Algorithm
var checksum *cobra.Command

func init() {

	ChecksumCmd().Flags().StringP("file", "f", "", "File to operate on")

	ChecksumCmd().Flags().Bool("md5", false, "checksum using md5")
	ChecksumCmd().Flags().Bool("sha1", false, "checksum using sha1")
	ChecksumCmd().Flags().Bool("sha256", false, "checksum using sha256")

	ChecksumCmd().MarkFlagRequired("file")

	hashFunctions = map[string]Algorithm{
		"md5":    GetMD5(),
		"sha1":   GetSHA1(),
		"sha256": GetSHA256(),
	}

	rootCmd.AddCommand(ChecksumCmd())
}

func GetChecksum(value []byte, hashFn string) string {
	return hashFunctions[hashFn].GetChecksum(value)
}

func ChecksumRun(cmd *cobra.Command, args []string) error {

	filename, _ := cmd.Flags().GetString("file")

	fs := File(filename).CheckFile()

	if fs.Err != nil {
		fmt.Println(fs.Err)
		return fs.Err
	} else {

		algorithmFlag := ""
		contents, _ := ioutil.ReadFile(filename)

		for _, algorithm := range hashFunctions {
			match, err := cmd.Flags().GetBool(algorithm.GetType())
			if err != nil {
				return err
			}
			if match {
				algorithmFlag = algorithm.GetType()
			}
		}

		fmt.Println(GetChecksum(contents, algorithmFlag))
	}

	return nil
}

func ChecksumCmd() *cobra.Command {
	if checksum == nil {
		checksum = &cobra.Command{
			Use:   "checksum",
			Short: "Generate checksum of a file",
			Long: `Checksum retrieves the cryptographic hash of a file.
		A valid text file and an algorithm flag (MD5, SHA1, SHA256) are needed.`,
			RunE: ChecksumRun,
		}
	}
	return checksum
}
