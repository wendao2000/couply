package config

import (
	"errors"
	"os"
	"strconv"
)

func InitConfig() (*Config, error) {
	redisDB, err := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64)
	if err != nil {
		return nil, errors.New("invalid env: REDIS_DB")
	}

	return &Config{
		AuthJWTIssuer: os.Getenv("AUTH_JWT_ISSUER"),
		AuthJWTSecret: os.Getenv("AUTH_JWT_SECRET"),
		AuthPassSalt:  os.Getenv("AUTH_PASS_SALT"),

		DSN: os.Getenv("DSN"),

		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       int(redisDB),
	}, nil
}
