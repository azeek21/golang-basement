// note this only works if x11 is available
package ex02

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lX11
// #include <stdlib.h>
// #include "application.h"
import "C"

func CreateWindow(title string) {
	width := C.uint(300)
	height := C.uint(200)
	x := C.int(500)
	y := C.int(500)
	nTitle := C.CString(title)
	C.MakeWindowWithTitle(x, y, width, height, nTitle)
}
