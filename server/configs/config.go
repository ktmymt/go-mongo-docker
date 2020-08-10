package configs

import "os"

// Config object
type Config struct {
	Env     string        `env:"ENV"`
	MongoDB MongoDBConfig `json:"mongodb"`
	Host    string        `env:"APP_HOST"`
	Port    string        `env:"APP_PORT"`
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:     os.Getenv("ENV"),
		MongoDB: GetMongoDBConfig(),
		Host:    os.Getenv("APP_HOST"),
		Port:    os.Getenv("APP_PORT"),
	}
}
