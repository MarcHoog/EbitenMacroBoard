package windows_api

import (
	"fmt"
	"unsafe"
)

var (
	procSendInput = user32.NewProc("SendInput")
)

func SendTextUniversal() {
	type keyboardInput struct {
		wVk         uint16
		wScan       uint16
		dwFlags     uint32
		time        uint32
		dwExtraInfo uint64
	}

	type input struct {
		inputType uint32
		ki        keyboardInput
		padding   uint64
	}

	var i input
	i.inputType = 1 //INPUT_KEYBOARD
	i.ki.wVk = 0x41 // virtual key code for a
	ret, _, err := procSendInput.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		uintptr(unsafe.Sizeof(i)),
	)
	fmt.Printf("ret: %v error: %v", ret, err)

}
