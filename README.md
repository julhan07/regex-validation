# Utils for Regex Validation

## How to use

```
import (
	"fmt"

	regexvalidation "github.com/julhan07/regex-validation/email"
)

func main() {
	email := "kampasi.dev@email.com"
	emailStatus := regexvalidation.ValidateEmail(email)
	fmt.Println(fmt.Sprintf("Your email %s : status %s", email, emailStatus))
}


```
