package web

import (
	"net/http"
)

func (s *service) postSignIn(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, signInPath, http.StatusSeeOther)
}
