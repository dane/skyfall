package v1

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/dane/skyfall/proto/gen/go/service/v1"
)

var (
	rxName            = regexp.MustCompile(`\A[a-z_]+\z`)
	rxAlphaLower      = regexp.MustCompile(`[a-z]+`)
	rxAlphaUpper      = regexp.MustCompile(`[A-Z]+`)
	rxNumeric         = regexp.MustCompile(`[0-9]+`)
	rxNonAlphaNumeric = regexp.MustCompile(`[^A-Za-z0-9]+`)
)

type Validator interface {
	CreateAccount(*pb.CreateAccountRequest) error
	UpdateAccount(*pb.UpdateAccountRequest) error
}

type validator struct{}

func NewValidator() Validator {
	return validator{}
}

func (v validator) CreateAccount(req *pb.CreateAccountRequest) error {
	return wrap(codes.InvalidArgument,
		validation.ValidateStruct(req,
			validation.Field(&req.Name, nameRules()...),
			validation.Field(&req.Password,
				validation.Required,
				validation.Length(8, 100),
				validation.NewStringRule(passwordRequirements, "must meet requirements"),
			),
			validation.Field(&req.PasswordConfirmation,
				validation.Required,
				validation.NewStringRule(passwordMatch(req.Password), "must match password"),
			),
		),
	)
}

func (v validator) UpdateAccount(req *pb.UpdateAccountRequest) error {
	return wrap(codes.InvalidArgument,
		validation.ValidateStruct(req,
			validation.Field(&req.Name, validation.When(req.Name != "", nameRules()...)),
		),
	)
}

func wrap(code codes.Code, err error) error {
	if err != nil {
		return status.Error(code, err.Error())
	}

	return nil
}

func nameRules() []validation.Rule {
	return []validation.Rule{
		validation.Required,
		validation.Length(3, 20),
		validation.Match(rxName),
	}
}

func passwordRequirements(password string) bool {
	var matchCount int
	if rxAlphaLower.MatchString(password) {
		matchCount++
	}

	if rxAlphaUpper.MatchString(password) {
		matchCount++
	}

	if rxNumeric.MatchString(password) {
		matchCount++
	}

	if rxNonAlphaNumeric.MatchString(password) {
		matchCount++
	}

	return matchCount >= 3
}

func passwordMatch(password string) func(string) bool {
	return func(confirmation string) bool {
		return password == confirmation
	}
}
