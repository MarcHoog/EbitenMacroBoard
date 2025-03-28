package windows_api

import (
	"fmt"
	"golang.org/x/sys/windows"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	getForegroundWindow = user32.NewProc("GetForegroundWindow")
	findWindowW         = user32.NewProc("FindWindowW")
	sendMessageW        = user32.NewProc("SendMessageW")
	setWindowLongProc   = user32.NewProc("SetWindowLongW")
	getWindowLongProc   = user32.NewProc("GetWindowLongW")
	activeWindowHandle  atomic.Value
)

const (
	GWL_EXSTYLE      = uintptr(^uint32(19)) // TODO: Read up on this lol
	WS_EX_NOACTIVATE = 0x08000000           // Prevents the window from being activated
)

func loadWindowHandle() (uintptr, error) {

	err := fmt.Errorf("could not load window handle")

	val := activeWindowHandle.Load()
	if val == nil {
		return 0, err
	}

	hwnd, ok := val.(uintptr)
	if !ok || hwnd == 0 {
		return 0, err
	}

	return hwnd, nil
}

func getActiveWindowHandle() uintptr {
	hwnd, _, _ := getForegroundWindow.Call()
	return hwnd
}

func findWindowByTitle(title string) uintptr {
	titleUTF16, _ := syscall.UTF16PtrFromString(title)
	hwnd, _, _ := findWindowW.Call(0, uintptr(unsafe.Pointer(titleUTF16)))
	return hwnd
}

func GetCurrentWindowTitle() string {
	hwnd, _ := loadWindowHandle()

	if windows.IsWindow(windows.HWND(hwnd)) == false {
		return ""
	}

	buf := make([]uint16, 256)
	sendMessageW.Call(hwnd, 0x000D, uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

func TrackActiveWindow() {
	var ownHwnd uintptr
	for ownHwnd == 0 {
		ownHwnd = findWindowByTitle("MacroBoard")
		if ownHwnd == 0 {
			fmt.Println("Waiting for MacroBoard window...")
			time.Sleep(500 * time.Millisecond) // Adjust delay as needed
		}
	}

	fmt.Printf("Found MacroBoard Window: 0x%X\n", ownHwnd)
	currentStyle, _, err := getWindowLongProc.Call(ownHwnd, GWL_EXSTYLE)
	if err != nil {
		fmt.Printf("Getting Current Window Style: %v\n", err)
	}
	newStyle := currentStyle | WS_EX_NOACTIVATE
	_, _, err = setWindowLongProc.Call(ownHwnd, GWL_EXSTYLE, newStyle)
	if err != nil {
		fmt.Printf("Injecting WS NOactivate: %v\n", err)
	}

	for {
		hwnd := getActiveWindowHandle()
		if hwnd != ownHwnd {
			activeWindowHandle.Store(hwnd)
			fmt.Printf("Stored Window ID: 0x%X\n", hwnd)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
