package config

type config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"prod"` // Struct tags
	StoragePath string     `yaml:"storage_path" env-requied:"true"`
	Server      HttpServer `yaml:"http_server"`
}
type HttpServer struct {
	Address string
}
