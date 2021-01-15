// +build darwin

package notification

import (
	"bytes"
	"fmt"
	"os/exec"
)

var (
	commandHandler = "osascript"
	flag           = "-e"
)

// Display function creates a notification
// using a system call and dependent on
// the notification structure
func (content *Content) Display() error {
	commandString := fmt.Sprintf(`display notification "%s" with title "%s"`, content.Message, content.Title)
	execCommand := exec.Command(commandHandler, flag, commandString)
	var output bytes.Buffer
	var stderr bytes.Buffer
	execCommand.Stdout = &output
	execCommand.Stderr = &stderr
	err := execCommand.Run()
	if err != nil {
		return err
	}
	return nil
}
