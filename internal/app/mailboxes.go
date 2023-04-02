package app

import (
	"log"

	imap "github.com/emersion/go-imap"
	client "github.com/emersion/go-imap/client"
)

func ListMailboxes(serverName string) {
	serverCredentials := mustFindCredentials(serverName)
	c := createClient(serverCredentials)

	// Don't forget to logout
	defer c.Logout()

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}

func CreateMailbox(serverName string, mailboxNames ...string) {
	serverCredentials := mustFindCredentials(serverName)
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(serverCredentials.DialString(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(serverCredentials.Login, serverCredentials.Password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	log.Println("Creating mailboxes")
	for _, mailboxName := range mailboxNames {
		if err := c.Create(mailboxName); err != nil {
			log.Printf("Unable to create mailbox %s: %s\n", mailboxName, err)
		}
	}
	log.Println("Mailboxes was created sucessfully")
}
