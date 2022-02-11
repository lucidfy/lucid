package env

import (
	"os"

	"github.com/daison12006013/gorvel/internal/facade/path"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	basepath := path.PathTo("/")
	LoadEnvFrom(basepath)
}

func LoadEnvFrom(basepath string) {
	LoadFile(basepath + ".env")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	LoadFile(basepath + ".env." + env)
}

func LoadFile(filepath string) {
	err := godotenv.Load(filepath)
	if err != nil {
		panic(err)
	}
}
