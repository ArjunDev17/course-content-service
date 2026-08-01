package config

import (
	"fmt"
	"os"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
}

type AppConfig struct {
	Name string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type KafkaConfig struct {
	Brokers []string
}

func Load() *Config {

	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "course-content-service"),
			Port: getEnv("APP_PORT", "8080"),
		},

		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5433"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "course_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},

		Kafka: KafkaConfig{
			Brokers: []string{
				getEnv("KAFKA_BROKER", "localhost:9092"),
			},
		},
	}
}

func (d DatabaseConfig) DSN() string {

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.Name,
		d.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
