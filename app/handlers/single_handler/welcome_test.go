package single_handler

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
)

func init() {
	env.LoadEnvForTests()
}

func TestWelcome(t *testing.T) {
	rt := WelcomeRoute.TestLoad(nil)
	rt.Uses(t)
	rt.AssertStatus(200)
	rt.AssertResponseContains(`Welcome to <a href="https://github.com/lucidfy/lucid" class="font-light">Lucid</a>`)
}
