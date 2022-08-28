package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv("MONGOURI")
}

func EnvGrpcPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv("GRPCPORT")
}
