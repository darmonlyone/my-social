package postgres

import "github.com/caarlos0/env/v6"

var postgresConfigPrefix = "POSTGRES_"

type Config struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD,unset" envDefault:"nopass"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	DBName   string `env:"DB_NAME"`
	SSL      string `env:"SSL"`
}

func NewConfigFromENV() (*Config, error) {
	opts := env.Options{
		Prefix:          postgresConfigPrefix,
		RequiredIfNoDef: true,
	}
	var c Config
	if err := env.Parse(&c, opts); err != nil {
		return nil, err
	}
	return &c, nil
}
