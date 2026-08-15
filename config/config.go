package config

import (
	"os"
)

type Config struct {
	Port  string
	DBDSN string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbDsn := os.Getenv("DB_DSN")
	if dbDsn == "" {
		dbDsn = "postgres://myuser:mypassword@localhost:5432/userdb?sslmode=disable"
	}

	return Config{
		Port:  port,
		DBDSN: dbDsn,
	}
}
