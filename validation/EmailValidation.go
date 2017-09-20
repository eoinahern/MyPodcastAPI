package validation

import "strings"

type EmailValidation struct {
}

func (e *EmailValidation) CheckEmailValid(email string) bool {

	if len(email) > 10 && strings.Contains(email, "@") && strings.Contains(email, ".") {
		return true
	}

	return false

}
