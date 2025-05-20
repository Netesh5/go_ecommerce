package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"` // Struct tags
	StoragePath string     `yaml:"storage_path" env-requied:"true"`
	Server      HttpServer `yaml:"http_server"`
	ApiVersion  string     `yaml:"api_version" env:"API_VERSION" env-default:"v1"`
}
type HttpServer struct {
	Address string
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

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Printf("Config loaded successfully: %s", config.Env)
	return &config

}
