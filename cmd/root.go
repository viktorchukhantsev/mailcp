package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mailcp",
	Short: "mailcp is an emails transfer tool",
	Long:  `mailcp is an emails transfer tool. The purpose of mailcp is to migrate emails or to backup emails.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
