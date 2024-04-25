package env

import (
	"os"
	"strconv"
)

var Configs = dotEnv{
	Pg: getPgConfig(),
}

type PgConfig struct {
	User     string
	Password string

	Host string
	Port int16

	Database string
}

type dotEnv struct {
	Pg PgConfig
}

func getPgConfig() PgConfig {
	port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 16)
	if err != nil {
		port = 5432
	}

	return PgConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     int16(port),
		Database: os.Getenv("DB_NAME"),
	}
}
