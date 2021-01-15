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
	var commandString string
	if content.IsAlert {
		commandString = fmt.Sprintf(`display alert "%s" message "%s"`, content.Title, content.Message)
	} else {
		commandString = fmt.Sprintf(`display notification "%s" with title "%s"`, content.Message, content.Title)
	}
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

//support for the say command coming soon
//osascript -e 'say "Hello world"'

//add support for icon path
