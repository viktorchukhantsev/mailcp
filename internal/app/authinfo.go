package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"mailcp/internal/helpers"
)

type Machine struct {
	Name     string
	Login    string
	Port     string
	Password string
}

func (m Machine) Valid() bool {
	if len(m.Password) == 0 {
		return false
	}

	if len(m.Name) == 0 {
		return false
	}

	if len(m.Port) == 0 {
		return false
	}

	if len(m.Login) == 0 {
		return false
	}

	return true
}

func (m Machine) DialString() string {
	return fmt.Sprintf("%s:%s", m.Name, m.PortNumber())
}

func (m Machine) PortNumber() string {
	if _, err := strconv.ParseInt(m.Port, 10, 64); err == nil {
		return m.Port
	}

	switch m.Port {
	case "imap":
		return "143"
	case "pop3":
		return "110"
	case "imaps":
		return "993"
	case "pop3s":
		return "995"
	}

	return "993"
}

type Authinfo struct {
	Machines []Machine
}

func ParseAuthinfo() Authinfo {
	authinfo, err := helpers.DecryptAuthinfo()
	if err != nil {
		log.Fatalf("Unable to decrypt ~/.authinfo.gpg")
	}

	res := Authinfo{}

	for _, line := range strings.Split(authinfo, "\n") {
		if []rune(line)[0] == []rune("#")[0] {
			continue
		}
		machine := Machine{}
		splitedLine := strings.Split(line, " ")
		for i := 0; i <= len(splitedLine)-2; i += 2 {
			switch op, val := splitedLine[i], splitedLine[i+1]; op {
			case "machine":
				machine.Name = strings.TrimSpace(val)
			case "login":
				machine.Login = strings.TrimSpace(val)
			case "port":
				machine.Port = strings.TrimSpace(val)
			case "password":
				machine.Password = helpers.RemoveQuotes(strings.TrimSpace(val))
			}
		}
		res.Machines = append(res.Machines, machine)
	}

	return res
}
