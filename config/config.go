package config

type Config struct {
	DatabaseDSN   string
	Auth0Audience string
	Auth0Issuer   string
	AUTH0JWKSURL  string
	AllowOrigins  []string
}
