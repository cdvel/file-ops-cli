package cmd

import(
    "fmt"

    "github.com/spf13/cobra"
)


func init(){
    rootCmd.AddCommand(checksumCmd)
}

var checksumCmd = &cobra.Command{
        Use: "checksum",
        Short: "Generate checksum of a file",
   		Long:  `Checksum retrieves the cryptographic hash of a file.
A valid text file and an algorithm flag (MD5, SHA1, SHA256) are needed.`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Unimplemented")
        },
}
