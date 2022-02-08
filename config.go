package regexvalidation

import "regexp"

var (
	EMAIL_FORMAT = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	PHONE_FORMAT = regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,10}$`)
	URL_FORMAT   = regexp.MustCompile(`((http|https)://)(www.)?” 
+ “[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]” 
+ “{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)`)
)
