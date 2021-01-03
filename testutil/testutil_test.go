package testutil_test

import (
	"regexp"
	"testing"

	"github.com/dane/skyfall/testutil"
)

func TestNewString(t *testing.T) {
	rx := regexp.MustCompile(`\A[a-z]+\z`)
	for i := 1; i < 50; i++ {
		s := testutil.NewString(t, i)
		if got := len(s); got != i {
			t.Errorf("got length %d; want %d; value %q", got, i, s)
		}

		if !rx.MatchString(s) {
			t.Errorf("value %q did not match regex", s)
		}
	}
}
