package windows_api

import (
	"fmt"
	"unsafe"
)

var (
	procSendInput = user32.NewProc("SendInput")
)

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

func SendTextUniversal(r rune) {

	var ks []input

	if r <= 0xFFFF {
		ks = append(ks, input{
			inputType: 1,
			ki: keyboardInput{
				wVk:     0,
				wScan:   uint16(r),
				dwFlags: 0x0004,
			},
		})

		ks = append(ks, input{
			inputType: 1,
			ki: keyboardInput{
				wVk:     0,
				wScan:   uint16(r),
				dwFlags: 0x0004 | 0x0002,
			},
		})
	} else {
		// Subtract 0x10000 to get the normalized code point
		codepoint := uint32(r) - 0x10000
		// High surrogate (top 10 bits of the code point + 0xD800)
		highSurrogate := uint16((codepoint >> 10) + 0xD800)
		// Low surrogate (bottom 10 bits of the code point + 0xDC00)
		lowSurrogate := uint16((codepoint & 0x3FF) + 0xDC00)

		// Add the high surrogate
		ks = append(ks, input{
			inputType: 1,
			ki: keyboardInput{
				wVk:     0,
				wScan:   highSurrogate,
				dwFlags: 0x0004,
			},
		})

		// Add the low surrogate
		ks = append(ks, input{
			inputType: 1,
			ki: keyboardInput{
				wVk:     0,
				wScan:   lowSurrogate,
				dwFlags: 0x0004,
			},
		})

		ks = append(ks, input{
			inputType: 1,
			ki: keyboardInput{
				wVk:     0,
				wScan:   highSurrogate,
				dwFlags: 0x0004 | 0x0002,
			},
		})

		ks = append(ks, input{
			inputType: 1,
			ki: keyboardInput{
				wVk:     0,
				wScan:   lowSurrogate,
				dwFlags: 0x0004 | 0x0002,
			},
		})
	}

	ret, _, err := procSendInput.Call(
		uintptr(len(ks)),
		uintptr(unsafe.Pointer(&ks[0])),
		uintptr(unsafe.Sizeof(input{})),
	)
	if ret != uintptr(len(ks)) {
		fmt.Printf("Warning: Only %d out of %d inputs were sent!\n", ret, len(ks))
	} else {
		fmt.Println("All inputs were successfully sent.")
	}

	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Printf("Error: %v\n", err)
	}

}
