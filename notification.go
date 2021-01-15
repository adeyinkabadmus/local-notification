package notification

//Content is a structure
//for the notification
type Content struct {
	Title    string
	Message  string
	IsAlert  bool
	RunAfter float32
}

//Fire function creates a notificatin
//content struct and fires a notification
//shell commands
func Fire(title string, message string, isAlert bool, runAfter float32) {
	content := &Content{
		Title:   title,
		Message: message,
		IsAlert: isAlert,
	}
	content.Display()
}
