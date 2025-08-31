package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type MongoConfig struct {
	URI               string `mapstructure:"uri"`
	Database          string `mapstructure:"database"`
	CoursesCollection string `mapstructure:"courses_collection"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Mongo  MongoConfig  `mapstructure:"mongo"`
}

var Cfg Config

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("mongo.uri", "mongodb://localhost:27017")
	viper.SetDefault("mongo.database", "edtech")
	viper.SetDefault("mongo.courses_collection", "courses")

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// basic debug log
	log.Printf("Config loaded (server port=%d, mongo.uri=%s, mongo.db=%s)",
		Cfg.Server.Port, Cfg.Mongo.URI, Cfg.Mongo.Database)

	// give some time for logs to flush in dev
	time.Sleep(10 * time.Millisecond)
}
