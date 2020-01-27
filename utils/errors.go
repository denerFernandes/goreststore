package utils

import (
	"errors"
	"strings"
)

// Error - Format error
func Error(err string) error {

	if strings.Contains(err, "name") {
		return errors.New("Name Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "password") {
		return errors.New("Incorrect Password")
	}

	return errors.New("Incorrect Details")

}
