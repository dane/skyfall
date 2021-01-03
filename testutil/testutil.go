package testutil

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func NewString(t *testing.T, n int) string {
	var s string
	b := make([]byte, 1)
	for i := 0; i < n; i++ {
		alpha := rand.Intn(25) + int('a')
		b[0] = byte(alpha)
		s = fmt.Sprintf("%s%s", s, b)
	}
	return s
}
