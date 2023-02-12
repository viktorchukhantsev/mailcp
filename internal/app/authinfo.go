package app

import (
	"fmt"
	"log"
	"strings"

	"mailcp/internal/helpers"
)

type Machine struct {
	Name     string
	Login    string
	Port     string
	Password string
}

type Authinfo struct {
	Machines []Machine
}

func ParseAuthinfo() *Authinfo {
	authinfo, err := helpers.DecryptAuthinfo()
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatalf("Unable to decrypt ~/.authinfo.gpg")
	}

	res := &Authinfo{}

	for _, line := range strings.Split(authinfo, "\n") {
		machine := Machine{}
		r := strings.NewReader(line)
		fmt.Fscanf(r, "machine %s login %s port %s password %s", &machine.Name, &machine.Login, &machine.Port, &machine.Password)
		machine.Password = helpers.RemoveQuotes(machine.Password)
		res.Machines = append(res.Machines, machine)
	}

	return res
}
