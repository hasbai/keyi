package app

import (
	"github.com/golang-jwt/jwt/v4"
	"keyi/auth"
	"testing"
	"time"
)

var claims = &auth.MyClaims{
	UID:          1,
	Permission:   auth.PUser,
	TenantID:     1,
	TenantAreaID: 1,
	Type:         "access",
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer:    "Issuer",
		ExpiresAt: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	},
}
var c = make(map[string]any)

func init() {
	c["claims"] = claims
	c["uid"] = 1
}

func BenchmarkClaims(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = c["claims"].(*auth.MyClaims)
	}
}

func BenchmarkUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = c["uid"].(int)
	}
}
