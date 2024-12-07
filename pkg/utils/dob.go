package utils

import (
	"time"
)

func ParseDOB(dob string) (time.Time, error) {
	return time.Parse(time.DateOnly, dob)
}
