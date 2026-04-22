package config

import "os"

type Config struct {
	DBHOST string
	DBPORT string
	DBUSER string
	DBPASS string
	DBNAME string
}

func Load() *Config {
	return &Config{
		DBHOST: getEnv("POSTGRES_HOST", "crypto-postgres"),
		DBPORT: getEnv("POSTGRES_PORT", "5432"),
		DBUSER: getEnv("POSTGRES_USER", "postgres_user"),
		DBPASS: getEnv("POSTGRES_PASSWORD", "postgres_password"),
		DBNAME: getEnv("POSTGRES_DB", "mydb"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

// package config

// import (
// 	"fmt"

// 	"github.com/kelseyhightower/envconfig"
// )

// type Config struct {
// 	DBHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
// 	DBPort     string `envconfig:"POSTGRES_PORT" default:"5432"`
// 	DBUser     string `envconfig:"POSTGRES_USER" default:"postgres_user"`
// 	DBPassword string `envconfig:"POSTGRES_PASSWORD" default:"postgres_password"`
// 	DBName     string `envconfig:"POSTGRES_DB" default:"mydb"`
// }

// func NewConfig() (*Config, error) {
// 	var cfg Config
// 	if err := envconfig.Process("", &cfg); err != nil {
// 		return nil, fmt.Errorf("procces envconfig: %w", err)
// 	}
// 	return &cfg, nil
// }

// func NewConfigMust() *Config {
// 	config, err := NewConfig()
// 	if err != nil {
// 		err = fmt.Errorf("Get config^ %w", err)
// 		panic(err)
// 	}
// 	return config
// }
