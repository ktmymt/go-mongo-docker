package configs

import (
	"os"
)

// MongoDBConfig object
type MongoDBConfig struct {
	URI string `env:"MONGO_URI"`
}

// GetMongoDBConfig でMngoDBConfigを受け取る
func GetMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI: os.Getenv("MONGO_URI"),
	}
}
