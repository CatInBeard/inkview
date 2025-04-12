package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview

extern void c_keyboard_handler(char *);
*/
import "C"
import (
	"unsafe"
)

//export keyboardHandler
func keyboardHandler(text *C.char) {
	userKeyboardHandler(C.GoString(text))
}

type KeyboardHandler func(string)

var userKeyboardHandler KeyboardHandler

func SetKeyboardHandler(kh KeyboardHandler) {
	userKeyboardHandler = kh
}

func OpenKeyboard(placeholder string, buflen int) {

	buffer := make([]byte, buflen)

	ctitle := C.CString(placeholder)
	defer C.free(unsafe.Pointer(ctitle))

	cbuffer := (*C.char)(unsafe.Pointer(&buffer[0]))

	var chandler C.iv_keyboardhandler
	chandler = (C.iv_keyboardhandler)(C.c_keyboard_handler)

	C.OpenKeyboard(ctitle, cbuffer, C.int(buflen), C.int(0), chandler)
}
