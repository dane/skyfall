package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidSignature = errors.New("invalid signature")
	ErrInvalidAlgorithm = errors.New("invalid algorithm")
)

const (
	Type = "JWT"
)

type Signer interface {
	Sign(string) string
	Verify(string) bool
	Algorithm() string
}

type Verifier func(interface{}) error

type Header struct {
	Type      string `json:"typ"`
	Algorithm string `json:"alg"`
}

func Parse(token string, signer Signer, verifier Verifier, dest interface{}) error {
	payload := strings.SplitN(token, ".", 3)
	if len(payload) != 3 {
		return ErrInvalidToken
	}

	reader := strings.NewReader(payload[0])
	dec := json.NewDecoder(base64.NewDecoder(base64.RawURLEncoding, reader))

	var header Header
	if err := dec.Decode(&header); err != nil {
		return fmt.Errorf("failed to decode header: %w", err)
	}

	if header.Algorithm != signer.Algorithm() {
		return ErrInvalidAlgorithm
	}

	// TODO: verify algorithm is a supported type

	if !signer.Verify(payload[1]) {
		return ErrInvalidSignature
	}

	reader.Reset(payload[1])
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("failed to decode payload: %w", err)
	}

	return verifier(dest)
}
