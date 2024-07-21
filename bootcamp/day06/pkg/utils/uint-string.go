package utils

import (
	"fmt"
	"strconv"
)

func UintToString(src uint) string {
	return fmt.Sprintf("%v", src)
}

func StringToUint(src string) (uint, error) {
	res64, err := strconv.ParseUint(src, 0, 64)
	if err != nil {
		return 0, err
	}
	return uint(res64), nil
}
