package regexvalidation

func ValidateEmail(email string) bool {

	if !EMAIL_FORMAT.MatchString(email) {
		return false
	}

	return true
}

func ValidatePhone(phone string) bool {

	if !PHONE_FORMAT.MatchString(phone) {
		return false
	}

	return true
}

func ValidateUrl(url string) bool {
	if !URL_FORMAT.MatchString(url) {
		return false
	}
	return true
}
