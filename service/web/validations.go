package web

import (
	"regexp"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	rxHasDigit       = regexp.MustCompile(`[0-9]+`)
	rxHasSpecialChar = regexp.MustCompile(`[-!@#$%^&*()_[\]{},.<>+=]+`)

	errPassword = validation.NewError("password", "must be a valid password")
)

const minPasswordMatch = 3

func usernameRules() []validation.Rule {
	return []validation.Rule{
		validation.Required,
		validation.Length(4, 20),
		is.Alphanumeric,
	}
}

func passwordRules() []validation.Rule {
	return []validation.Rule{
		validation.Required,
		validation.Length(10, 50),
		validation.By(func(v interface{}) error {
			password, err := validation.EnsureString(v)
			if err != nil {
				return err
			}

			var count int
			if govalidator.HasUpperCase(password) {
				count++
			}

			if govalidator.HasLowerCase(password) {
				count++
			}

			if rxHasDigit.MatchString(password) {
				count++
			}

			if rxHasSpecialChar.MatchString(password) {
				count++
			}

			if count < minPasswordMatch {
				return errPassword
			}

			return nil
		}),
	}
}
