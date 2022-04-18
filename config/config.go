package config

import (
	"fmt"
	"os"
)

type Config struct {
	Web
	Database
}

func GetConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}

	name := os.Getenv("DATABASE_NAME")
	if name == "" {
		name = "udv_test"
	}

	pass := os.Getenv("DATABASE_PASSWORD")
	if pass == "" {
		pass = "password"
	}

	usr := os.Getenv("DATABASE_USER")
	if usr == "" {
		usr = "udv_test"
	}

	fmt.Printf("user %s. password %s. port %s. host %s. \n", usr, pass, port, host)
	config := Config{
		Web: Web{
			port: port,
		},
		Database: Database{
			name:     name,
			password: pass,
			user:     usr,
			host:     host,
		},
	}

	return config
}
