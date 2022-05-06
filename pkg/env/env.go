package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lucidfy/lucid/pkg/facade/path"
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
