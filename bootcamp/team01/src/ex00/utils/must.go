package utils

import "log"

// NOTE: this stops process if any error
func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
