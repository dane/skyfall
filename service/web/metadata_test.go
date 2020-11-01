package web

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMetadata(t *testing.T) {
	tests := []struct {
		name string
		key  interface{}
		set  func(context.Context, Metadata) context.Context
		get  func(context.Context) (Metadata, error)
	}{
		{
			name: "incoming",
			key:  incomingCtxKey{},
			set:  SetIncomingMetadata,
			get:  GetIncomingMetadata,
		},
		{
			name: "outgoing",
			key:  outgoingCtxKey{},
			set:  SetOutgoingMetadata,
			get:  GetOutgoingMetadata,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			v := ctx.Value(tc.key)
			if v != nil {
				t.Fatal("expected nil")
			}

			want := Metadata{"test": "test"}
			ctx = tc.set(ctx, want)
			v = ctx.Value(tc.key)
			if v == nil {
				t.Fatal("expected not to be nil")
			}

			got, err := tc.get(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, want) {
				t.Fatal(cmp.Diff(got, want))
			}
		})
	}
}
