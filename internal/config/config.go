package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}

	configPath = filepath.Join(wd, configPath)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Config file does not exist:", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Cannot read config:", err)
	}

	return &cfg
}
