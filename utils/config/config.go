package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		MongoDatabase MongoDatabaseConfig
	}

	MongoDatabaseConfig struct {
		UseSRV   bool
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
	}
)

var appConfig Config

func GetAPPConfig() Config {
	// Load .env
	godotenv.Load()

	// Mongo database
	confUseSRV := os.Getenv("MONGO_USE_SRV")
	useSRV, _ := strconv.ParseBool(confUseSRV)
	appConfig.MongoDatabase.UseSRV = useSRV
	appConfig.MongoDatabase.Host = os.Getenv("MONGO_HOST")
	appConfig.MongoDatabase.Port = os.Getenv("MONGO_PORT")
	appConfig.MongoDatabase.Username = os.Getenv("MONGO_USERNAME")
	appConfig.MongoDatabase.Password = os.Getenv("MONGO_PASSWORD")
	appConfig.MongoDatabase.DBName = os.Getenv("MONGO_DBNAME")

	return appConfig
}
