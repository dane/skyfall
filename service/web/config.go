package web

import (
	"github.com/joeshaw/envdecode"
)

type Config struct {
	TemplatePath string `env:"SKYFALL_TEMPLATE_PATH,default=templates"`
	HTTPAddr     string `env:"SKYFALL_HTTP_ADDR,default=:8000"`
	CookieSecret string `env:"SKYFALL_COOKIE_SECRET,default=development"`
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	return c, envdecode.Decode(c)
}
