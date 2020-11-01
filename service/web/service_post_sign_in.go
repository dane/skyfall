package web

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (s *service) postSignIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(fieldNameUsername)
	password := r.FormValue(fieldNamePassword)
	if err := validatePostSignIn(username, password); err != nil {
		s.SetError(r, err)
		// set error message
		// redirect to GET /sign-in
		return
	}

	http.Redirect(w, r, signInPath, http.StatusSeeOther)
}

func validatePostSignIn(username, password string) error {
	return validation.Errors{
		"username": validation.Validate(&username, usernameRules()...),
		"password": validation.Validate(&password, passwordRules()...),
	}.Filter()
}
