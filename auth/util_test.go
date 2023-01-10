package auth

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	const id = 100
	code1, err := setCode(id)
	if err != nil {
		t.Error(err)
	}
	code2, err := getCode(id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(code1, code2)
	if code1 != code2 {
		t.Error("code1 != code2")
	}
}
