package v1

type Option interface {
	apply(*service)
}

func SetValidator(v Validator) Option {
	return setValidator{v}
}

type setValidator struct {
	validator Validator
}

func (v setValidator) apply(s *service) {
	s.validator = v.validator
}
