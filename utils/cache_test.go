package utils

import (
	"testing"
	"time"
)

func TestCacheString(t *testing.T) {
	var s string
	err := SetCache("test", "test", time.Second)
	if err != nil {
		t.Error(err)
	}

	err = GetCache("test", &s)
	if err != nil {
		t.Error(err)
	}
	if s != "test" {
		t.Error("cache string failed")
	}
}
