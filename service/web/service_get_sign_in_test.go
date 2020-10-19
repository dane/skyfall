package web

import (
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSignIn(t *testing.T) {
	mr := MockRender{T: t}
	s := &service{
		render: mr.Render,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/sign-in", nil)
	s.getSignIn(w, r)

	if got, want := mr.Name, "sessions/sign-in"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	data := s.Data(r)
	if !cmp.Equal(mr.Data, data) {
		t.Error(cmp.Diff(mr.Data, data))
	}
}
