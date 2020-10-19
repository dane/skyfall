package web

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.uber.org/zap"

	"github.com/dane/skyfall/jwt"
)

const (
	signInPath = "/sign-in"

	cookieUserSession = "_session"
)

type Service interface {
	Render(w http.ResponseWriter, name string, data Data) error
	Logger() *zap.Logger
	Data(r *http.Request) Data
}

type Data map[string]string

type Render func(w http.ResponseWriter, name string, data Data) error

type Handler func(s Service) http.Handler

type service struct {
	getSignInHandler  http.Handler
	postSignInHandler http.Handler

	errorInternalHandler         http.Handler
	errorMethodNotAllowedHandler http.Handler
	errorNotFoundHandler         http.Handler

	render Render
	logger *zap.Logger
	signer jwt.Signer
}

func New(cfg *Config, options ...Option) (http.Handler, error) {
	s := &service{}
	for _, opt := range options {
		opt.apply(s)
	}

	if err := validateService(s); err != nil {
		return nil, err
	}

	r := chi.NewRouter()
	r.Use(s.extractUserSession)
	r.Method(http.MethodGet, signInPath, s.getSignInHandler)
	r.Method(http.MethodPost, signInPath, s.postSignInHandler)

	return r, nil
}

func (s *service) Render(w http.ResponseWriter, name string, data Data) error {
	return s.render(w, name, data)
}

func (s *service) Logger() *zap.Logger {
	return s.logger
}

func (s *service) Data(r *http.Request) Data {
	data := make(Data)
	data["signInPath"] = signInPath

	return data
}

func (s *service) extractUserSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieUserSession)
		if err == http.ErrNoCookie {
			next.ServeHTTP(w, r)
			return
		}

		verifier := func(payload interface{}) error {
			session := payload.(*UserSession)
			now := time.Now().Unix()
			rules := []*validation.FieldRules{
				validation.Field(&session.ID, validation.Required, is.UUID),
				validation.Field(&session.Audience, validation.Required),
				validation.Field(&session.IssuedAt, validation.Required, validation.Max(now)),
				validation.Field(&session.NotBefore, validation.Required, validation.Max(now)),
				validation.Field(&session.ExpiresAt, validation.Required, validation.Min(now)),
			}

			if session.Subject != "" || session.Scope != nil {
				rules = append(rules,
					validation.Field(&session.Scope, validation.Required),
					validation.Field(&session.Subject, validation.Required, is.UUID),
				)
			}

			return validation.ValidateStruct(session, rules...)
		}

		var session UserSession
		if err := jwt.Parse(cookie.Value, s.signer, verifier, &session); err != nil {
			s.errorInternalHandler.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userSessionKey{}, session)
		next.ServeHTTP(w, r.Clone(ctx))
	})
}

func validateService(s *service) error {
	return validation.ValidateStruct(s,
		validation.Field(&s.getSignInHandler, validation.Required),
		validation.Field(&s.postSignInHandler, validation.Required),
		validation.Field(&s.logger, validation.Required),
		validation.Field(&s.render, validation.Required),
	)
}