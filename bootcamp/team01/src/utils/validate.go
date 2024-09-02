package utils

import (
	"errors"
	"fmt"
	"strconv"
)

const PORT_MIN = 8000
const PORT_MAX = 8999

var (
	SRV_ERR_BAD_PORT = errors.New(fmt.Sprintf("port is not valid. Must be between %d-%d", PORT_MIN, PORT_MAX))
)

// validates port and returns error if not valid containing the reason why validation didn't pass
func IsPortValid(port string) error {
	n, err := strconv.Atoi(port)
	if err != nil {
		return SRV_ERR_BAD_PORT
	}

	if n < PORT_MIN || n > PORT_MAX {
		return SRV_ERR_BAD_PORT
	}

	return nil
}

func IsValidUUID(str string) error {
	// TODO: implement uuid check
	return nil
}
