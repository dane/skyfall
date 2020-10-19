package web

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

func DefaultGetSignInHandler(s Service) http.Handler {
	return http.HandlerFunc(s.(*service).getSignIn)
}

func DefaultPostSignInHandler(s Service) http.Handler {
	return http.HandlerFunc(s.(*service).postSignIn)
}

func DefaultRender(templatePath string) Render {
	return func(w http.ResponseWriter, name string, data Data) error {
		fullPath := filepath.Join(templatePath, fmt.Sprintf("%s.mustache", name))
		tmpl, err := mustache.ParseFile(fullPath)
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		if err := tmpl.FRender(w, data); err != nil {
			return fmt.Errorf("failed to render template: %w", err)
		}

		return nil
	}
}
