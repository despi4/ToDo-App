package authsvc

import "time"

type JWTConfig struct {
	AccessSecret    string
	Refresh         string
	Issuer          string
	AccesTokenTTL   time.Duration
	RefreshTokenTTL time.Duration
}


