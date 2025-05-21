package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"` // Struct tags
	StoragePath string     `yaml:"storage_path" env-requied:"true"`
	Server      HttpServer `yaml:"http_server"`
	ApiVersion  string     `yaml:"api_version" env:"API_VERSION" env-default:"v1"`
	DbConfig    DBCOfig    `yaml:"db" env-required:"true"`
}
type HttpServer struct {
	Address string
}

type DBCOfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	User     string
	Password string
	DbName   string `yaml:"database" env-required:"true"`
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration files")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("config path is required")
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file not found: %s", configPath)
	}

	var config Config

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	config.DbConfig.User = os.Getenv("DB_USER")
	config.DbConfig.Password = os.Getenv("DB_PASSWORD")
	config.DbConfig.Host = os.Getenv("DB_HOST")
	config.DbConfig.DbName = os.Getenv("DB_NAME")

	if port, err := strconv.Atoi(os.Getenv("DB_PORT")); err == nil {
		config.DbConfig.Port = port
	} else {
		log.Fatalf("failed to convert DB_PORT to int: %v", err)
	}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Printf("Config loaded successfully: %s", config.Env)
	return &config

}
