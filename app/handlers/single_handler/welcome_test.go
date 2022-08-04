package single_handler

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
)

func init() {
	env.LoadEnvForTests()
}

func TestWelcome(t *testing.T) {
	rt := WelcomeRoute.LoadTester(nil, nil)
	rt.CallHandler()

	rt.AssertUsing(t)
	rt.AssertStatus(200)
	rt.AssertResponseContains(`Welcome to <a href="https://github.com/lucidfy/lucid" class="font-light">Lucid</a>`)
}

func TestWelcomeAsJson(t *testing.T) {
	rt := WelcomeRoute.LoadTester(nil, nil)
	rt.Request.Header.Set("Accept", "application/json")
	rt.CallHandler()

	rt.AssertUsing(t)
	rt.AssertStatus(200)
	rt.AssertResponseContains(`{"lang":{},"title":"Lucid Rocks!"}`)
}
