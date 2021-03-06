package v1_test

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"

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
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestUpdateAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.UpdateAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "name can be blank",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Name = ""
			},
			valid: true,
		},
		{
			name: "properties can be blank",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Properties = nil
			},
			valid: true,
		},
		{
			name: "name and properties can be blank",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Name = ""
				req.Properties = nil
			},
			valid: true,
		},
		{
			name:   "valid",
			modify: func(*pb.UpdateAccountRequest) {},
			valid:  true,
		},
		{
			name: "name is too short",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Name = testutil.NewString(t, 2)
			},
			message: "name: the length must be between 3 and 20.",
		},
		{
			name: "name is too long",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Name = testutil.NewString(t, 21)
			},
			message: "name: the length must be between 3 and 20.",
		},
		{
			name: "name cannot contain special characters",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Name = testutil.NewString(t, 10) + "!"
			},
			message: "name: must be in a valid format.",
		},
		{
			name: "name cannot contain dashes",
			modify: func(req *pb.UpdateAccountRequest) {
				req.Name = "example-name"
			},
			message: "name: must be in a valid format.",
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			properties, err := structpb.NewStruct(map[string]interface{}{
				"example-key": "example-value",
			})
			if err != nil {
				t.Fatal(err)
			}

			req := &pb.UpdateAccountRequest{
				Name:       "example",
				Properties: properties,
			}

			// Apply modififications to the Update request.
			tc.modify(req)

			err = validator.UpdateAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.DeleteAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "id cannot be blank",
			modify: func(req *pb.DeleteAccountRequest) {
				req.Id = ""
			},
			message: "id: cannot be blank.",
		},
		{
			name: "id cannot be numeric",
			modify: func(req *pb.DeleteAccountRequest) {
				req.Id = "123"
			},
			message: "id: must be a valid UUID.",
		},
		{
			name: "id cannot be alpha",
			modify: func(req *pb.DeleteAccountRequest) {
				req.Id = testutil.NewString(t, 10)
			},
			message: "id: must be a valid UUID.",
		},
		{
			name:   "valid",
			modify: func(*pb.DeleteAccountRequest) {},
			valid:  true,
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.DeleteAccountRequest{
				Id: uuid.New().String(),
			}

			// Apply modififications to the Delete request.
			tc.modify(req)

			err := validator.DeleteAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestVerifyAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.VerifyAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "id cannot be blank",
			modify: func(req *pb.VerifyAccountRequest) {
				req.Id = ""
			},
			message: "id: cannot be blank.",
		},
		{
			name: "id cannot be numeric",
			modify: func(req *pb.VerifyAccountRequest) {
				req.Id = "123"
			},
			message: "id: must be a valid UUID.",
		},
		{
			name: "id cannot be alpha",
			modify: func(req *pb.VerifyAccountRequest) {
				req.Id = testutil.NewString(t, 10)
			},
			message: "id: must be a valid UUID.",
		},
		{
			name:   "valid",
			modify: func(*pb.VerifyAccountRequest) {},
			valid:  true,
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.VerifyAccountRequest{
				Id: uuid.New().String(),
			}

			// Apply modififications to the Verify request.
			tc.modify(req)

			err := validator.VerifyAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestSuspendAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.SuspendAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "id cannot be blank",
			modify: func(req *pb.SuspendAccountRequest) {
				req.Id = ""
			},
			message: "id: cannot be blank.",
		},
		{
			name: "id cannot be numeric",
			modify: func(req *pb.SuspendAccountRequest) {
				req.Id = "123"
			},
			message: "id: must be a valid UUID.",
		},
		{
			name: "id cannot be alpha",
			modify: func(req *pb.SuspendAccountRequest) {
				req.Id = testutil.NewString(t, 10)
			},
			message: "id: must be a valid UUID.",
		},
		{
			name:   "valid",
			modify: func(*pb.SuspendAccountRequest) {},
			valid:  true,
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.SuspendAccountRequest{
				Id: uuid.New().String(),
			}

			// Apply modififications to the Suspend request.
			tc.modify(req)

			err := validator.SuspendAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestUndeleteAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.UndeleteAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "id cannot be blank",
			modify: func(req *pb.UndeleteAccountRequest) {
				req.Id = ""
			},
			message: "id: cannot be blank.",
		},
		{
			name: "id cannot be numeric",
			modify: func(req *pb.UndeleteAccountRequest) {
				req.Id = "123"
			},
			message: "id: must be a valid UUID.",
		},
		{
			name: "id cannot be alpha",
			modify: func(req *pb.UndeleteAccountRequest) {
				req.Id = testutil.NewString(t, 10)
			},
			message: "id: must be a valid UUID.",
		},
		{
			name:   "valid",
			modify: func(*pb.UndeleteAccountRequest) {},
			valid:  true,
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.UndeleteAccountRequest{
				Id: uuid.New().String(),
			}

			// Apply modififications to the Undelete request.
			tc.modify(req)

			err := validator.UndeleteAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestUnsuspendAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.UnsuspendAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "id cannot be blank",
			modify: func(req *pb.UnsuspendAccountRequest) {
				req.Id = ""
			},
			message: "id: cannot be blank.",
		},
		{
			name: "id cannot be numeric",
			modify: func(req *pb.UnsuspendAccountRequest) {
				req.Id = "123"
			},
			message: "id: must be a valid UUID.",
		},
		{
			name: "id cannot be alpha",
			modify: func(req *pb.UnsuspendAccountRequest) {
				req.Id = testutil.NewString(t, 10)
			},
			message: "id: must be a valid UUID.",
		},
		{
			name:   "valid",
			modify: func(*pb.UnsuspendAccountRequest) {},
			valid:  true,
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.UnsuspendAccountRequest{
				Id: uuid.New().String(),
			}

			// Apply modififications to the Unsuspend request.
			tc.modify(req)

			err := validator.UnsuspendAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestGetAccount(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.GetAccountRequest)
		valid   bool
		message string
	}{
		{
			name: "id cannot be blank",
			modify: func(req *pb.GetAccountRequest) {
				req.Id = ""
			},
			message: "id: cannot be blank.",
		},
		{
			name: "id cannot be numeric",
			modify: func(req *pb.GetAccountRequest) {
				req.Id = "123"
			},
			message: "id: must be a valid UUID.",
		},
		{
			name: "id cannot be alpha",
			modify: func(req *pb.GetAccountRequest) {
				req.Id = testutil.NewString(t, 10)
			},
			message: "id: must be a valid UUID.",
		},
		{
			name:   "valid",
			modify: func(*pb.GetAccountRequest) {},
			valid:  true,
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.GetAccountRequest{
				Id: uuid.New().String(),
			}

			// Apply modififications to the Get request.
			tc.modify(req)

			err := validator.GetAccount(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func TestGetAccountByName(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*pb.GetAccountByNameRequest)
		valid   bool
		message string
	}{
		{
			name:   "valid",
			modify: func(*pb.GetAccountByNameRequest) {},
			valid:  true,
		},
		{
			name: "missing name",
			modify: func(req *pb.GetAccountByNameRequest) {
				req.Name = ""
			},
			message: "name: cannot be blank.",
		},
		{
			name: "name is too short",
			modify: func(req *pb.GetAccountByNameRequest) {
				req.Name = testutil.NewString(t, 2)
			},
			message: "name: the length must be between 3 and 20.",
		},
		{
			name: "name is too long",
			modify: func(req *pb.GetAccountByNameRequest) {
				req.Name = testutil.NewString(t, 21)
			},
			message: "name: the length must be between 3 and 20.",
		},
		{
			name: "name cannot contain special characters",
			modify: func(req *pb.GetAccountByNameRequest) {
				req.Name = testutil.NewString(t, 10) + "!"
			},
			message: "name: must be in a valid format.",
		},
		{
			name: "name cannot contain dashes",
			modify: func(req *pb.GetAccountByNameRequest) {
				req.Name = "example-name"
			},
			message: "name: must be in a valid format.",
		},
	}

	validator := v1.NewValidator()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.GetAccountByNameRequest{
				Name: "example",
			}

			// Apply modififications to the create request.
			tc.modify(req)

			err := validator.GetAccountByName(req)
			testValidationError(t, tc.valid, tc.message, err)
		})
	}
}

func testValidationError(t *testing.T, valid bool, message string, err error) {
	t.Helper()
	if err == nil {
		if valid {
			return
		}
		t.Fatal("error was expected")
	} else {
		st := status.Convert(err)
		if got, want := st.Code(), codes.InvalidArgument; got != want {
			t.Errorf("got code %q; want %q", got, want)
		}

		if got := st.Message(); got != message {
			t.Errorf("got message %q; want %q", got, message)
		}
	}
}
