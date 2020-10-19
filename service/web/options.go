package web

import (
	"go.uber.org/zap"
)

type Option interface {
	apply(*service)
}

type applyer func(*service)

func (fn applyer) apply(s *service) {
	fn(s)
}

func SetGetSignInHandler(fn Handler) Option {
	return applyer(func(s *service) {
		s.getSignInHandler = fn(s)
	})
}

func SetPostSignInHandler(fn Handler) Option {
	return applyer(func(s *service) {
		s.postSignInHandler = fn(s)
	})
}

func SetRender(fn Render) Option {
	return applyer(func(s *service) {
		s.render = fn
	})
}

func SetLogger(logger *zap.Logger) Option {
	return applyer(func(s *service) {
		s.logger = logger
	})
}
