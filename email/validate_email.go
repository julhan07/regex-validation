package email

import (
	config "github.com/julhan07/regex-validation"
)

func ValidateEmail(email string) bool {

	if !config.EMAIL_FORMAT.MatchString(email) {
		return false
	}

	return true
}
