package utils

import (
	"fmt"
	"keyi/config"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println(MD5("123"))
}

func TestSendEmail(t *testing.T) {
	err := SendEmail([]string{config.Config.SmtpUser}, "test", "this is body")
	if err != nil {
		t.Error(err)
	}
}

func TestCache(t *testing.T) {
	const key = "key"
	const value = "value"

	err := SetCache(key, value, 0)
	if err != nil {
		t.Error(err)
	}

	var result string
	err = GetCache(key, &result)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	if value != result {
		t.Error("cache error", "value: ", value, "result: ", result)
	}
}
