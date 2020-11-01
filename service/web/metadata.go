package web

import (
	"context"
	"errors"
)

var ErrCastMetadata = errors.New("failed to cast metadata")

type Metadata map[string]string
type incomingCtxKey struct{}
type outgoingCtxKey struct{}

func SetIncomingMetadata(ctx context.Context, md Metadata) context.Context {
	return setMetadata(ctx, incomingCtxKey{}, md)
}

func SetOutgoingMetadata(ctx context.Context, md Metadata) context.Context {
	return setMetadata(ctx, outgoingCtxKey{}, md)
}

func GetIncomingMetadata(ctx context.Context) (Metadata, error) {
	return getMetadata(ctx, incomingCtxKey{})
}

func GetOutgoingMetadata(ctx context.Context) (Metadata, error) {
	return getMetadata(ctx, outgoingCtxKey{})
}

func setMetadata(ctx context.Context, lookup interface{}, md Metadata) context.Context {
	return context.WithValue(ctx, lookup, md)
}

func getMetadata(ctx context.Context, lookup interface{}) (Metadata, error) {
	v := ctx.Value(lookup)
	if v == nil {
		return Metadata{}, nil
	}

	md, ok := v.(Metadata)
	if !ok {
		return nil, ErrCastMetadata
	}

	return md, nil
}
