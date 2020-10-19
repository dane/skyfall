package web

import (
	"net/http/httptest"
	"testing"
)

func TestDefaultRender(t *testing.T) {
	w := httptest.NewRecorder()
	data := make(Data)
	data["test"] = "example"
	render := DefaultRender("../../testdata/templates")

	if err := render(w, "test", data); err != nil {
		t.Fatal(err)
	}

	if got, want := w.Body.String(), "<p>example</p>\n"; got != want {
		t.Fatalf("got %q; want %q", got, want)
	}
}

func TestDefaultRenderErrors(t *testing.T) {
	t.Run("when the templat path is wrong", func(t *testing.T) {
		w := httptest.NewRecorder()
		render := DefaultRender("invalid/path")
		err := render(w, "test", nil)
		if err == nil {
			t.Fatal("error expected")
		}
	})

	t.Run("when the template is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := make(Data)
		data["test"] = "example"
		render := DefaultRender("../../testdata/templates")

		err := render(w, "invalid", data)
		if err == nil {
			t.Fatal("error expected")
		}
	})
}
