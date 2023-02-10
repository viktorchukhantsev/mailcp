package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionNumber string
var gitHash string
var buildTime string
var author string
var email string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mailcp",
	Long:  `All software has versions. This is mailcp`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mailcp %s -- %s build at %s\n", versionNumber, gitHash, buildTime)
		fmt.Printf("Send feedback to %s\n", email)
		fmt.Printf("Best Regards, %s\n", author)
	},
}
