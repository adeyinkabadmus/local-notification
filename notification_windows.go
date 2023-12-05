// +build windows

package notification

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	messageBoxW          = user32.NewProc("MessageBoxW")
	getConsoleWindow     = user32.NewProc("GetConsoleWindow")
	shGetFolderPathW     = shell32.NewProc("SHGetFolderPathW")
	shell32              = syscall.NewLazyDLL("shell32.dll")
	shellExecuteExW      = shell32.NewProc("ShellExecuteExW")
	sW_SHOW              = uintptr(5)
	SEE_MASK_NOCLOSEPROCESS = uintptr(0x00000040)
)

func main() {
	title := "Notification Title"
	message := "Notification Message"
	iconPath := "C:\\Path\\To\\Your\\Logo.ico"

	showNotification(title, message, iconPath)
}

func showNotification(title, message, iconPath string) {
	consoleWindow, _, _ := getConsoleWindow.Call()

	if consoleWindow == 0 {
		// If running as a console application, use MessageBox
		messageBoxW.Call(0,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
			0)
	} else {
		// If not a console application, use ShellExecuteEx to display notification
		iconPathPtr, _ := syscall.UTF16PtrFromString(iconPath)
		params := &shellExecuteExInfo{
			cbSize: sizeofShellExecuteExInfo,
			fMask:  SEE_MASK_NOCLOSEPROCESS,
			hwnd:   0,
			lpVerb: nil,
			lpFile: nil,
			lpParameters: nil,
			lpDirectory:  nil,
			nShow:        sW_SHOW,
			hInstApp:     0,
			lpIDList:     nil,
			lpClass:      nil,
			hkeyClass:    0,
			dwHotKey:     0,
			hIconOrMonitor:   0,
			hProcess:     0,
		}

		success, _, _ := shellExecuteExW.Call(uintptr(unsafe.Pointer(params)))

		if success == 1 {
			fmt.Println("Notification displayed successfully.")
		} else {
			fmt.Println("Error displaying notification.")
		}
	}
}

const (
	sizeofShellExecuteExInfo = unsafe.Sizeof(shellExecuteExInfo{})
)

type shellExecuteExInfo struct {
	cbSize          uintptr
	fMask           uintptr
	hwnd            uintptr
	lpVerb          *uint16
	lpFile          *uint16
	lpParameters    *uint16
	lpDirectory     *uint16
	nShow           uintptr
	hInstApp        uintptr
	lpIDList        *uint16
	lpClass         *uint16
	hkeyClass       uintptr
	dwHotKey        uintptr
	hIconOrMonitor  uintptr
	hProcess        uintptr
}
