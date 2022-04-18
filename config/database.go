package config

import "fmt"

type Database struct {
	user     string
	password string
	name     string
	host     string
}

func (cfg Database) DSN() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cfg.host, cfg.user, cfg.name, cfg.password)
}
