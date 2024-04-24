package config

import "time"

var (
	JWT_SECRET     = "secret"
	JWT_EXPIRATION = time.Hour * 24
)
