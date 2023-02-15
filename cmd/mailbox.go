package cmd

import (
	"github.com/spf13/cobra"

	"mailcp/internal/app"
)

func init() {
	mailboxCmd.AddCommand(mailboxLsCmd)
	mailboxCmd.AddCommand(mailboxCreateCmd)
	rootCmd.AddCommand(mailboxCmd)
}

var mailboxCmd = &cobra.Command{
	TraverseChildren: true,
	Use:              "mailbox",
	Short:            "Mailboxes commands",
	Long:             `Mailboxes commands. It uses ~/.authinfo.gpg and GPG to fetch server credentials`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var mailboxLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List mailboxes",
	Long:  "List mailboxes available on server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.ListMailboxes(args[0])
	},
}

var mailboxCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create mailboxes",
	Long:  "Create mailboxes on server",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		app.CreateMailbox(args[0], args[1:]...)
	},
}
