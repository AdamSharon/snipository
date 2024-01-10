package history

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetLatestFromHistory(historyFileEnvVar string) (string, error) {
	historyFilePath, historyFileFound := os.LookupEnv(historyFileEnvVar)
	if !historyFileFound {
		return "", fmt.Errorf("HISTFILE env var not found, please make sure you're using zsh")
	}

	getLatestCommandBash := fmt.Sprintf("echo $(cat %s | tail -n 2 | head -n 1)", historyFilePath)

	// why not use GO default way to read file? because this file can be huge,
	// and reading all + reverse it will be slow. this is actually faster.
	out, err := exec.Command("/bin/sh", "-c",
		getLatestCommandBash).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", err)
		return "", err
	}

	var command string
	commandUnparsed := strings.TrimSpace(string(out))
	commandParts := strings.SplitN(commandUnparsed, ";", 2) // Split into 2 parts on space

	// If there's a command part after the number
	if len(commandParts) > 1 {
		command = commandParts[1]
	} else {
		return "", fmt.Errorf("no command found in history")
	}

	return handleCommandEscapeCharacters(command), nil
}

// handleCommandEscapeCharacters handles escape characters in the command
// currently only handles new lines
func handleCommandEscapeCharacters(command string) string {
	// Replace all new lines with \n so it will be escaped
	command = strings.ReplaceAll(command, "\n", "\\n")

	return command
}
