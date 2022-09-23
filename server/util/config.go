package util

import "github.com/joho/godotenv"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

var Current *Config

func LoadConfig() {
	godotenv.Load()

	Current = &Config{
		DB: &DBConfig{
			Name:     MustString("DB_NAME", "shawty"),
			Username: MustString("DB_USERNAME", "postgres"),
			Password: MustString("DB_PASSWORD", "1234"),
			Host:     MustString("DB_HOST", "localhost"),
			Port:     MustString("DB_PORT", "5432"),
		},
	}

}
