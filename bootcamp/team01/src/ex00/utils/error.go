package utils

import (
	"errors"
	"fmt"
)

func WithPrefix(prefix string, err error) error {
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %s", prefix, err.Error()))
	}
	return err
}
