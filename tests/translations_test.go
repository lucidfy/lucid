package tests

import (
	"os"
	"testing"

	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/resources/translations"
)

func TestTranslateWithNoEnv(t *testing.T) {
	got := translations.Get("validations.email", helpers.MS{
		":field": "Email",
	})
	expect := "Email is not a valid email address!"

	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func TestTranslateUsingChineseSimplified(t *testing.T) {
	os.Setenv("APP_LANGUAGE", "zh-CN")

	got := translations.Get("validations.email", helpers.MS{
		":field": "Email",
	})
	expect := "Email 不是有效的电子邮件地址"

	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func TestTranslateUsingChineseTraditional(t *testing.T) {
	os.Setenv("APP_LANGUAGE", "zh-TW")

	got := translations.Get("validations.email", helpers.MS{
		":field": "Email",
	})
	expect := "Email 不是有效的電子郵件地址"

	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}
