package config

type Config struct {
	AuthPassSalt  string
	AuthJWTSecret string

	AuthJWTIssuer string
	DSN           string

	RedisAddr     string
	RedisPassword string
	RedisDB       int
}
