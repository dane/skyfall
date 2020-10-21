package web

import (
	"strings"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func TestUsernameRules(t *testing.T) {
	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "blank username",
			value: "",
			valid: false,
		},
		{
			name:  "too short",
			value: "dh",
			valid: false,
		},
		{
			name:  "too long",
			value: "12345678901234567890a", // 21 characters
			valid: false,
		},
		{
			name:  "not alphanumeric",
			value: "test!",
			valid: false,
		},
		{
			name:  "alphanumeric",
			value: "test",
			valid: true,
		},
		{
			name:  "alphanumeric",
			value: "test1234", // 8 characters
			valid: true,
		},
	}

	rules := usernameRules()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validation.Validate(tc.value, rules...)
			if tc.valid && err != nil {
				t.Fatal(err)
			}

			if !tc.valid && err == nil {
				t.Fatal("error wanted")
			}
		})
	}
}

func TestPasswordRules(t *testing.T) {
	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "blank password",
			value: "",
			valid: false,
		},
		{
			name:  "too short",
			value: "Ab1@test!", // 9 characters
			valid: false,
		},
		{
			name:  "too long",
			value: "Ab1$" + strings.Repeat("a", 50), // 54 characters
			valid: false,
		},
		{
			name:  "uppercase and lowercase only",
			value: "TestTestTe",
			valid: false,
		},
		{
			name:  "lowercase and digit only",
			value: "test111122",
			valid: false,
		},
		{
			name:  "lowercase and non-alphanumeric only",
			value: "test!@#$",
			valid: false,
		},
		{
			name:  "uppercase, lowercase, digit, and non-alphanumeric",
			value: "Test1234!!",
			valid: true,
		},
		{
			name:  "uppercase, lowercase, and digit",
			value: "Test123456",
			valid: true,
		},
	}

	rules := passwordRules()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validation.Validate(tc.value, rules...)
			if tc.valid && err != nil {
				t.Fatal(err)
			}

			if !tc.valid && err == nil {
				t.Fatal("error wanted")
			}
		})
	}
}
