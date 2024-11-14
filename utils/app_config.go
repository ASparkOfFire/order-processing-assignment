package utils

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		DatabaseName string `validate:"required"`
		DatabaseHost string `validate:"required"`
		DatabasePort string `validate:"required"`
		DatabaseUser string `validate:"required"`
		DatabasePass string `validate:"required"`
	}
	Server struct {
		HTTPListenAddress string `validate:"required"`
		HTTPPort          int    `validate:"required,min=1024,max=65535"`
	}
}

var AppConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	AppConfig.Server.HTTPListenAddress = os.Getenv("HTTP_LISTEN_ADDRESS")
	AppConfig.Server.HTTPPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	AppConfig.Database.DatabaseHost = os.Getenv("DATABASE_HOST")
	AppConfig.Database.DatabasePort = os.Getenv("DATABASE_PORT")
	AppConfig.Database.DatabaseName = os.Getenv("DATABASE_NAME")
	AppConfig.Database.DatabaseUser = os.Getenv("DATABASE_USER")
	AppConfig.Database.DatabasePass = os.Getenv("DATABASE_PASS")

	// validate the environment variables
	validationInstance := validator.New(validator.WithRequiredStructEnabled())
	if err := validationInstance.Struct(&AppConfig); err != nil {
		log.Fatalf("err loading environment variables: %v\n", err)
	}
}
