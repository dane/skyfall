package web

import (
	"net/http"

	"go.uber.org/zap"
)

func (s *service) getSignIn(w http.ResponseWriter, r *http.Request) {
	if err := s.Render(w, "sessions/sign-in", s.Data(r)); err != nil {
		s.Logger().Error("failed to render GET sign-in", zap.Error(err))
	}
}
