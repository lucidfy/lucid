package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/daison12006013/gorvel/pkg/env"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
)

func main() {
	env.LoadEnv()
	url := os.Getenv("SCHEMA") + "://" + os.Getenv("HOST") + ":" + os.Getenv("PORT")
	logger.Info("Serving at " + url)
	openbrowser(url)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		logger.Fatal(err)
	}
}
