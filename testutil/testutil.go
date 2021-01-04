package testutil

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	pbv1 "github.com/dane/skyfall/proto/gen/go/service/v1"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func NewString(t *testing.T, n int) string {
	t.Helper()

	var s string
	b := make([]byte, 1)
	for i := 0; i < n; i++ {
		alpha := rand.Intn(25) + int('a')
		b[0] = byte(alpha)
		s = fmt.Sprintf("%s%s", s, b)
	}
	return s
}

func IgnoreFields(t *testing.T) []cmp.Option {
	t.Helper()

	return []cmp.Option{
		cmpopts.IgnoreUnexported(pbv1.CreateAccountResponse{}, pbv1.Account{}),
		cmpopts.IgnoreFields(pbv1.Account{}, "Id", "CreatedAt", "UpdatedAt"),
	}
}
