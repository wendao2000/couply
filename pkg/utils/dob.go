package utils

import (
	"time"
)

func ParseDOB(dob string) (string, error) {
	if len(dob) == 0 {
		return "", nil
	}
	_, err := time.Parse(time.DateOnly, dob)
	if err != nil {
		return "", err
	}
	return dob, nil
}
