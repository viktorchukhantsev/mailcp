package app

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func ListMailboxes(serverName string) {
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

func CopyMailbox(fromServerName, origMailboxName, toServerName, destMailboxName string) {
	fromServerCredentials := mustFindCredentials(fromServerName)
	toServerCredentials := mustFindCredentials(toServerName)
	log.Println("From: Connecting to server...")

	// Connect to server
	fromClient, err := client.DialTLS(fromServerCredentials.DialString(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("From: Connected")

	// Don't forget to logout
	defer fromClient.Logout()

	// Login
	if err := fromClient.Login(fromServerCredentials.Login, fromServerCredentials.Password); err != nil {
		log.Fatal(err)
	}
	log.Println("From: Logged in")

	log.Println("To: Connecting to server...")
	// Connect to server
	toClient, err := client.DialTLS(toServerCredentials.DialString(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("To: Connected")

	// Don't forget to logout
	defer toClient.Logout()

	// Login
	if err := toClient.Login(toServerCredentials.Login, toServerCredentials.Password); err != nil {
		log.Fatal(err)
	}
	log.Println("To: Logged in")
}
