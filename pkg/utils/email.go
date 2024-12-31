package utils

import (
	"regexp"
)

const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var emailRegex = regexp.MustCompile(emailRegexPattern)

// ValidateEmail checks if the email format is valid
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}
