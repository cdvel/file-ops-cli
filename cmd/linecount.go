package cmd

import(
    "fmt"

    "github.com/spf13/cobra"
)


func init(){
    rootCmd.AddCommand(linecountCmd)
}

var linecountCmd = &cobra.Command{
        Use: "linecount",
        Short: "Count lines in a file",
   		Long:  `Linecount counts the number of lines in a text file. A valid text file is required.`,

        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Unimplemented")
        },
}
