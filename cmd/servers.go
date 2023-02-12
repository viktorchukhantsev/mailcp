package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"mailcp/internal/app"
)

func init() {
	rootCmd.AddCommand(serversCmd)
}

var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "Print servers available for mailcp",
	Long:  `It uses ~/.authinfo.gpg and GPG`,
	Run: func(cmd *cobra.Command, args []string) {
		authInfo := app.ParseAuthinfo()

		for _, machine := range authInfo.Machines {
			fmt.Printf("- %s\n", machine.Name)
			fmt.Printf("  %s\n", machine.Port)
			fmt.Printf("  %s\n", machine.Login)
			fmt.Printf("  %s\n", strings.Repeat("*", len(machine.Password)))
		}
	},
}
