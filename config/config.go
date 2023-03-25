package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPAddr   string `yaml:"httpAddr"`
	PgURL      string `env:"PG_URL"`
	FileSystem string `yaml:"fileSystem"`
}

func New() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
