package auth

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	user := User{TenantID: 1, TenantAreaID: 1}
	access, refresh, err := GenerateTokens(&user)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(access)
	fmt.Println(refresh)

	claims, err := ParseToken(access)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(claims)
}
