package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer HTTPServerConfig
	Logger     LoggerConfig
	Mongo      MongoConfig
}

type MongoConfig struct {
	URI    string `env:"MONGO_URI" envDefault:"mongodb://mongo:mongo@localhost:27017"`
	DBName string `env:"MONGO_DB_NAME" envDefault:"tanca-event-mongo"`
}

type HTTPServerConfig struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Mode string `env:"MODE" envDefault:"development"`
}

type LoggerConfig struct {
	Level    string `env:"LOGGER_LEVEL" envDefault:"debug"`
	Mode     string `env:"MODE" envDefault:"development"`
	Encoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

// Load loads the configuration from the environment variables.
// dòng này chạy thứ 2  Load cấu hình
func Load() (*Config, error) {
	// Load .env file
	//dòng này chạy thứ 2.1 để load các biến môi trường từ file .env
	_ = godotenv.Load()

	var config Config
	//dòng này chạy thứ 2.2 để phân tích các biến môi trường và gán chúng vào cấu trúc Config và return về main.go
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
