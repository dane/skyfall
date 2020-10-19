package web

type userSessionKey struct{}

type UserSession struct {
	ID        string   `json:"jti"`
	Subject   string   `json:"sub,omitempty"`
	Scope     []string `json:"scope,omitempty"`
	Error     string   `json:"err,omitempty"`
	Audience  string   `json:"aud"`
	IssuedAt  int64    `json:"iat"`
	NotBefore int64    `json:"nbf"`
	ExpiresAt int64    `json:"exp"`
}
