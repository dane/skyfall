package v1_test

import (
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/dane/skyfall/proto/gen/go/service/v1"
	"github.com/dane/skyfall/service/v1"
	"github.com/dane/skyfall/testutil"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.CreateAccountRequest)
		valid   bool
		message string
	}{
		{
			name:   "valid",
			modify: func(*pb.CreateAccountRequest) {},
			valid:  true,
		},
		{
			name: "missing name",
			modify: func(req *pb.CreateAccountRequest) {
				req.Name = ""
			},
			message: "name: cannot be blank.",
		},
		{
			name: "name is too short",
			modify: func(req *pb.CreateAccountRequest) {
				req.Name = testutil.NewString(t, 2)
			},
			message: "name: the length must be between 3 and 20.",
		},
		{
			name: "name is too long",
			modify: func(req *pb.CreateAccountRequest) {
				req.Name = testutil.NewString(t, 21)
			},
			message: "name: the length must be between 3 and 20.",
		},
		{
			name: "name cannot contain special characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Name = testutil.NewString(t, 10) + "!"
			},
			message: "name: must be in a valid format.",
		},
		{
			name: "name cannot contain dashes",
			modify: func(req *pb.CreateAccountRequest) {
				req.Name = "example-name"
			},
			message: "name: must be in a valid format.",
		},
		{
			name: "password only contains alpha lower characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = testutil.NewString(t, 8)
				req.PasswordConfirmation = req.Password
			},
			message: "password: must meet requirements.",
		},
		{
			name: "password only contains alpha upper characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = strings.ToUpper(testutil.NewString(t, 8))
				req.PasswordConfirmation = req.Password
			},
			message: "password: must meet requirements.",
		},
		{
			name: "password only contains numeric characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = "12345678"
				req.PasswordConfirmation = req.Password
			},
			message: "password: must meet requirements.",
		},
		{
			name: "password only contains special characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = "!@#$%^&*"
				req.PasswordConfirmation = req.Password
			},
			message: "password: must meet requirements.",
		},
		{
			name: "password only contains alpha-numeric characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = testutil.NewString(t, 8) + "1"
				req.PasswordConfirmation = req.Password
			},
			message: "password: must meet requirements.",
		},
		{
			name: "password only contains alpha and special characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = testutil.NewString(t, 8) + "!"
				req.PasswordConfirmation = req.Password
			},
			message: "password: must meet requirements.",
		},
		{
			name: "password contains alpha lower, numeric, and special characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = testutil.NewString(t, 8) + "!2"
				req.PasswordConfirmation = req.Password
			},
			valid: true,
		},
		{
			name: "password contains alpha upper, numeric, and special characters",
			modify: func(req *pb.CreateAccountRequest) {
				req.Password = strings.ToUpper(testutil.NewString(t, 8)) + "!2"
				req.PasswordConfirmation = req.Password
			},
			valid: true,
		},
		{
			name: "password confirmation does not match password",
			modify: func(req *pb.CreateAccountRequest) {
				req.PasswordConfirmation = "bad"
			},
			message: "password_confirmation: must match password.",
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.CreateAccountRequest{
				Name:                 "example",
				Password:             "Example!",
				PasswordConfirmation: "Example!",
			}

			// Apply modififications to the create request.
			tc.modify(req)

			err := validator.CreateAccount(req)
			if err == nil {
				if tc.valid {
					return
				}
				t.Fatal("error was expected")
			} else {
				st := status.Convert(err)
				if got, want := st.Code(), codes.InvalidArgument; got != want {
					t.Errorf("got code %q; want %q", got, want)
				}

				if got := st.Message(); got != tc.message {
					t.Errorf("got message %q; want %q", got, tc.message)
				}
			}
		})
	}
}
