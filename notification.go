package notification

import (
	"runtime"
)

//Content is a structure
//for the notification
type Content struct {
	Title    string
	Message  string
	IsAlert  bool
	RunAfter float32
}

func main() {
	underlyingOs := runtime.GOOS
	if underlyingOs == "windows" {
		//select action
	} else if underlyingOs == "darwin" {
		//select another action
	}
}

//Fire does fire mate
func Fire(title string, message string, isAlert bool, runAfter float32) {
	content := &Content{
		Title:   title,
		Message: message,
		IsAlert: isAlert,
	}
	content.Display()
}
