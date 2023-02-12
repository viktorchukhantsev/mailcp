package helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const gpgArgs = " -q --for-your-eyes-only --no-tty -d $HOME/.authinfo.gpg"

var gpgCommands = [...]string{"gpg", "gpg2"}

func DecryptAuthinfo() (string, error) {
	gpgCommand, err := findGPG()

	if err != nil {
		return "", err
	}

	home := os.Getenv("HOME")
	if len(home) == 0 {
		return "", errors.New("Unable to obtain home directory")
	}

	authinfo, err := getGPGOutput(gpgCommand + strings.Replace(gpgArgs, "$HOME", home, 1))
	if err != nil {
		return "", fmt.Errorf("Unable to decrypt $HOME/.authinfo.gpg: %v", err)
	}

	return authinfo, nil
}
