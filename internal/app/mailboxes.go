package app

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func ListMailboxes(serverName string) {
	authInfo := ParseAuthinfo()
	var serverCredentials Machine
	var found bool

	for i := range authInfo.Machines {
		if authInfo.Machines[i].Name == serverName {
			serverCredentials = authInfo.Machines[i]
			found = true
			break
		}
	}
	if !found {
		log.Fatal("Unable to found this server in authinfo")
	}

	if serverCredentials.Valid() {
		log.Fatalf("%s credentials is invalid\n", serverName)
	}

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
