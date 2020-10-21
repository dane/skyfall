package web

import (
	"net/http"

	"go.uber.org/zap"
)

func (s *service) getSignIn(w http.ResponseWriter, r *http.Request) {
	data := s.Data(r)
	data["usernameFieldName"] = fieldNameUsername
	data["passwordFieldName"] = fieldNamePassword

	if err := s.Render(w, "sessions/sign-in", data); err != nil {
		s.Logger().Error("failed to render GET sign-in", zap.Error(err))
	}
}
