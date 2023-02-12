package helpers

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	which "github.com/hairyhenderson/go-which"
)

func findGPG() (string, error) {
	for _, gpgCommand := range gpgCommands {
		if which.Found(gpgCommand) {
			return gpgCommand, nil
		}
	}

	return "", errors.New("Unable to find installed GPG")
}

func getGPGOutput(command string) (string, error) {
	var err error
	var stdOutFile *os.File = os.Stdout
	reader, writer := io.Pipe()
	scannerStopped := make(chan struct{})
	var secureOutput []string
	stdOut := bufio.NewWriter(stdOutFile)

	go func() {
		defer close(scannerStopped)

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output := scanner.Text()
			if strings.Contains(output, "machine") {
				secureOutput = append(secureOutput, output)
			} else {
				stdOut.WriteString(output)
			}
		}
	}()

	cmdArgs := strings.Fields(command)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = writer
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	go func() {
		_ = cmd.Wait()
		writer.Close()
	}()

	<-scannerStopped

	return strings.Join(secureOutput, "\n"), nil
}
