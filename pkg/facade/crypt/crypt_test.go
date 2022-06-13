package crypt

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
)

func init() {
	env.LoadEnv()
}

func TestCrypt(t *testing.T) {
	text := "1234qwer"
	encrypted, err := Encrypt(text)

	if err != nil {
		t.Errorf("encrypt is not working %q", err)
	}

	decrypted, err := Decrypt(encrypted)

	if err != nil {
		t.Errorf("decrypt is not working %q", err)
	}

	if decrypted != text {
		t.Errorf("got %q, wanted %q", decrypted, text)
	}
}
