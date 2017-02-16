package concurrent

import (
	"bytes"
	"os/exec"
	"strings"
)

// ExecuteCommand executes the command passed to the application
func ExecuteCommand(command string) (string, error) {
	//Parse the command
	commandArray := strings.Split(command, " ")
	cmdName := commandArray[0]
	cmdArgs := commandArray[1:]
	cmd := exec.Command(cmdName, cmdArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
