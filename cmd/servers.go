package cmd

import (
	"log"
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
		log.Println("Listing available servers")
		for _, machine := range app.ParseAuthinfo().Machines {
			log.Printf("- %s\n", machine.Name)
			log.Printf("  %s\n", machine.PortNumber())
			log.Printf("  %s\n", machine.Login)
			log.Printf("  %s\n", strings.Repeat("*", len(machine.Password)))
		}
	},
}
