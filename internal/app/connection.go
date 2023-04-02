package app

import (
	"errors"
	"log"

	client "github.com/emersion/go-imap/client"
	sasl "github.com/emersion/go-sasl"
)

func createClient(serverCredentials Machine) *client.Client {
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(serverCredentials.DialString(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	saslClient := sasl.NewLoginClient(serverCredentials.Login, serverCredentials.Password)

	// Login
	if err := c.Login(serverCredentials.Login, serverCredentials.Password); err != nil {
		if errors.Is(client.ErrLoginDisabled, err) {
			if err := c.Authenticate(saslClient); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	log.Println("Logged in")
	// Select Mailbox

	return c
}
