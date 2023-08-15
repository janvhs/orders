package config

import "github.com/caarlos0/env/v6"

type DB struct {
	DSN string `env:"DB_DSN"`
}

type Server struct {
	Host string `env:"SERVER_HOST"`
	Port int    `env:"SERVER_PORT"`
}

type LDAP struct {
	Host         string `env:"LDAP_HOST"`
	Port         int    `env:"LDAP_PORT"`
	BaseDN       string `env:"LDAP_BASE_DN"`
	BindDN       string `env:"LDAP_BIND_DN,unset"`
	BindPassword string `env:"LDAP_BIND_PASSWORD,unset"`
}

type Config struct {
	IsDevelopment bool `env:"DEV"`
	DB            DB
	Server        Server
	LDAP          LDAP
}

func NewFromEnv() (Config, error) {
	cfg := Default()

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
