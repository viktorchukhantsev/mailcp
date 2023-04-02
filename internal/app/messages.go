package app

import (
	"log"
)

func ContentMailbox(serverName, mailboxName string) {
	serverCredentials := mustFindCredentials(serverName)
	c := createClient(serverCredentials)
	// Don't forget to logout
	defer c.Logout()

	mailboxStatus, err := c.Select(mailboxName, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Mailbox contains %d messages\n", mailboxStatus.Messages)
}
