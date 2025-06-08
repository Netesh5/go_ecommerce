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
	Env              string     `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"` // Struct tags
	StoragePath      string     `yaml:"storage_path" env-requied:"true"`
	Server           HttpServer `yaml:"http_server"`
	ApiVersion       string     `yaml:"api_version" env:"API_VERSION" env-default:"v1"`
	DbConfig         DBCOfig    `yaml:"db" env-required:"true"`
	OtpConfig        OTPConfig  `env-required:"true"`
	CloudinaryConfig Cloudinary `env-required:"true"`
}
type HttpServer struct {
	Address string
}

type DBCOfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type OTPConfig struct {
	AccountSID string
	AuthToken  string
	ServiceSID string
}

type Cloudinary struct {
	CloudName string
	APIKey    string
	APISecret string
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

	// OPT

	config.OtpConfig.AccountSID = os.Getenv("SENDGRID_ACCOUNT_SID")
	config.OtpConfig.AuthToken = os.Getenv("SENDGRID_AUTH_TOKEN")
	config.OtpConfig.ServiceSID = os.Getenv("SENDGRID_SERVICE_SID")

	// Cloudinary
	config.CloudinaryConfig.CloudName = os.Getenv("CLOUDINARY_NAME")
	config.CloudinaryConfig.APIKey = os.Getenv("CLOUDINARY_API_KEY")
	config.CloudinaryConfig.APISecret = os.Getenv("CLOUDINARY_API_SECRET")

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
