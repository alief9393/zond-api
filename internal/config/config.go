package config

import "os"

type Config struct {
	PostgresConn string
}

func LoadConfig() (Config, error) {
	return Config{
		PostgresConn: os.Getenv("POSTGRES_CONN"),
	}, nil
}
