package ex00

import (
	"errors"
	"unsafe"
)

var ERR_OUT_OF_BOUNDS = errors.New("Given index is out of bounds of provided slice")
var ERR_EMPTY_SLICE = errors.New("Provided slice is empty")
var ERR_INVALID_INDEX = errors.New("Give index is invalid. Meaning it's either non numberic or negative")

func GetElement(arr []int, idx int) (int, error) {
	res := 0
	if idx < 0 {
		return res, ERR_INVALID_INDEX
	}

	arrLen := len(arr)
	if arrLen == 0 {
		return res, ERR_EMPTY_SLICE
	}

	if idx >= arrLen {
		return res, ERR_OUT_OF_BOUNDS
	}

	ptrToStart := unsafe.Pointer(&arr[0])
	elemSize := unsafe.Sizeof(int(0))
	res = *(*int)(unsafe.Add(ptrToStart, elemSize*uintptr(idx)))
	return res, nil
}
