package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/daison12006013/gorvel/pkg/env"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/urls"
)

func main() {
	var u string

	env.LoadEnv()
	flag.StringVar(&u, "url", urls.BaseUrl(nil), "URL to Open")
	flag.Parse()

	logger.Info("Serving at " + u)
	openBrowser(u)
}

func openBrowser(url string) {
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
