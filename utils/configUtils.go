package utils

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var envLoaded bool = false

func LoadEnv() error {
	if envLoaded == true {
		return nil
	}

	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		pathFile, err := filepath.Abs(os.Getenv("GOPATH") + "/src/busca-cep-go/.env")
		if err != nil {
			return err
		}
		envFile = pathFile
	}

	err := godotenv.Load(envFile)
	if err != nil {
		return errors.New("error_loading_env_file")
	}

	envLoaded = true
	return nil
}
