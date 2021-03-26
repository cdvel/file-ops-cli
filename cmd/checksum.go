package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	. "fops-cli/filesys"
	"github.com/spf13/cobra"
	"hash"
	"io/ioutil"
)

func init() {
	fmt.Println(hashAlgorithms)

	checksumCmd.Flags().StringP("file", "f", "", "File to operate on")

	checksumCmd.Flags().Bool("md5", false, "checksum using md5")
	checksumCmd.Flags().Bool("sha1", false, "checksum using sha1")
	checksumCmd.Flags().Bool("sha256", false, "checksum using sha256")

	checksumCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(checksumCmd)
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

		is_md5, _ := cmd.Flags().GetBool("md5")
		is_sha1, _ := cmd.Flags().GetBool("sha1")
		is_sha256, _ := cmd.Flags().GetBool("sha256")

		if is_md5 {
			fmt.Println(fmt.Sprintf("%x", md5.Sum(contents)))
		}

		if is_sha1 {
			fmt.Println(fmt.Sprintf("%x", sha1.Sum(contents)))
		}

		if is_sha256 {
			fmt.Println(fmt.Sprintf("%x", sha256.Sum256(contents)))
		}

	},
}
