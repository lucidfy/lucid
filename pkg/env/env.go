package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lucidfy/lucid/pkg/facade/logger"
	"github.com/lucidfy/lucid/pkg/facade/path"
)

func LoadEnv() {
	basepath := path.PathTo("/")
	LoadEnvFrom(basepath)
}

func LoadEnvForTests() {
	os.Setenv("LUCID_TESTS", "1")
	LoadEnv()
}

func LoadEnvFrom(basepath string) {
	LoadFile(basepath + ".env")
	env := os.Getenv("APP_ENV")
	if env != "" {
		LoadFile(basepath + ".env." + env)
	}
}

func LoadFile(filepath string) {
	err := godotenv.Load(filepath)
	if err != nil {
		logger.Error("Cannot load .env file", err)
	}
}
