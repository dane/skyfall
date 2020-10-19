package web

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	config := &Config{}
	options := []Option{
		SetGetSignInHandler(DefaultGetSignInHandler),
		SetPostSignInHandler(DefaultPostSignInHandler),
		SetRender(DefaultRender("test")),
		SetLogger(zap.NewNop()),
	}

	t.Run("with all required options", func(t *testing.T) {
		handler, err := New(config, options...)
		if handler == nil {
			t.Error("handler expected")
		}

		if err != nil {
			t.Errorf("got %q; want %v", err, nil)
		}
	})

	for skip, opt := range options {
		t.Run(fmt.Sprintf("without %v", opt), func(t *testing.T) {
			var incomplete []Option
			for i, option := range options {
				if i == skip {
					continue
				}

				incomplete = append(incomplete, option)
			}

			_, err := New(config, incomplete...)
			if err == nil {
				t.Fatal("error expected")
			}
		})
	}
}
