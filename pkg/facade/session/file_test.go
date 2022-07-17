package session

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
)

func init() {
	env.LoadEnvForTests()
}

func TestFileSessionFlash(t *testing.T) {
	ses := &FileSession{
		SessionKey: "golangtest",
		FileMode:   0744,
	}

	ses.PutFlashMap("data", map[string]interface{}{
		"key1": true,
		"key2": "key2 is string",
		"key3": 100.99,
		"key4": 1,
	})

	data := *ses.GetFlashMap("data")

	if data["key1"] != true {
		t.Errorf("Flash is not working for [key1]")
	}

	if data["key2"] != "key2 is string" {
		t.Errorf("Flash is not working for [key2]")
	}

	if data["key3"] != 100.99 {
		t.Errorf("Flash is not working for [key3]")
	}

	if data["key4"].(float64) != 1 {
		t.Errorf("Flash is not working for [key4]")
	}
}
