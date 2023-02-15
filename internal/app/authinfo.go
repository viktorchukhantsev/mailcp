package app

import (
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
		log.Fatalf("Unable to decrypt ~/.authinfo.gpg")
	}

	res := &Authinfo{}

	for _, line := range strings.Split(authinfo, "\n") {
		if []rune(line)[0] == []rune("#")[0] {
			continue
		}
		machine := Machine{}
		splitedLine := strings.Split(line, " ")
		for i := 0; i <= len(splitedLine)-2; i += 2 {
			switch op, val := splitedLine[i], splitedLine[i+1]; op {
			case "machine":
				machine.Name = val
			case "login":
				machine.Login = val
			case "port":
				machine.Port = val
			case "password":
				machine.Password = helpers.RemoveQuotes(val)
			}
		}
		res.Machines = append(res.Machines, machine)
	}

	return res
}
