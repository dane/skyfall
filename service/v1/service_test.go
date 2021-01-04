package v1_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	pb "github.com/dane/skyfall/proto/gen/go/service/v1"
	"github.com/dane/skyfall/service/v1"
	"github.com/dane/skyfall/testutil"
)

func TestServiceCreateAccount(t *testing.T) {
	// Define context and request.
	ctx := context.Background()
	req := &pb.CreateAccountRequest{
		Name:     testutil.NewString(t, 5),
		Password: "Example1",
	}
	req.PasswordConfirmation = req.Password

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Define validator and service.
	v := v1.NewMockValidator(ctrl)
	v.EXPECT().CreateAccount(req)
	svc, err := v1.New(v1.SetValidator(v))
	if err != nil {
		t.Fatal(err)
	}

	// Issue request.
	got, err := svc.CreateAccount(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	want := &pb.CreateAccountResponse{
		Account: &pb.Account{
			Name: req.Name,
		},
	}

	if !cmp.Equal(got, want, testutil.IgnoreFields(t)...) {
		t.Fatal(cmp.Diff(got, want, testutil.IgnoreFields(t)...))
	}

}
