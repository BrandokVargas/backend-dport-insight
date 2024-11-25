package main

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}
	return nil
}

func validateEnvironments() error {
	if strings.TrimSpace(os.Getenv("SERVER_PORT")) == "" {
		return errors.New("the SERVER_PORT env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS")) == "" {
		return errors.New("the ALLOWED_ORIGINS env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_METHODS")) == "" {
		return errors.New("the ALLOWED_METHODS env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("IMAGES_DIR")) == "" {
		return errors.New("the IMAGES_DIR env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("JWT_SECRET_KEY")) == "" {
		return errors.New("the JWT_SECRET_KEY env is mandatory")
	}

	//Data integrated
	if strings.TrimSpace(os.Getenv("DB_USER_DB_INTEGRATED")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_PASSWORD_DB_INTEGRATED")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_HOST_DB_INTEGRATED")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_PORT_DB_INTEGRATED")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_NAME_DB_INTEGRATED")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_SSL_MODE_DB_INTEGRATED")) == "" {
		return errors.New("the env is mandatory")
	}

	return nil
}
