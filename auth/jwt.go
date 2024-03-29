package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"keyi/config"
	. "keyi/utils"
	"time"
)

type TokenType = string

type TokenTypeStruct struct {
	TokenType
	Expires time.Duration
}

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

var (
	AccessToken = TokenTypeStruct{
		TokenTypeAccess,
		time.Hour * 2,
	}
	RefreshToken = TokenTypeStruct{
		TokenTypeRefresh,
		time.Hour * 24 * 7,
	}
)

type MyClaims struct {
	jwt.RegisteredClaims
	UID          int        `json:"uid"` // user id
	Permission   Permission `json:"permission"`
	TenantID     int        `json:"tenant_id"`
	TenantAreaID int        `json:"tenant_area_id"`
	Type         string     `json:"type"`
}

type TokenInfo struct {
	UserID       int        `json:"user_id"`
	Permission   Permission `json:"permission"`
	TenantID     int        `json:"tenant_id"`
	TenantAreaID int        `json:"tenant_area_id"`
}

type HasTokenInfo interface {
	GetTokenInfo() *TokenInfo
}

func (t *MyClaims) GetTokenInfo() *TokenInfo {
	return &TokenInfo{
		UserID:       t.UID,
		Permission:   t.Permission,
		TenantID:     t.TenantID,
		TenantAreaID: t.TenantAreaID,
	}
}

func (t *MyClaims) GetPermission() Permission {
	return t.Permission
}

func generateToken(info *TokenInfo, tokenType TokenTypeStruct) (string, error) {
	claim := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenType.Expires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.Config.SiteName,
		},
		Type:         tokenType.TokenType,
		TenantID:     info.TenantID,
		TenantAreaID: info.TenantAreaID,
		Permission:   info.Permission,
		UID:          info.UserID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.Config.SecretKey))
}

func GenerateTokens(obj HasTokenInfo) (string, string, error) {
	info := obj.GetTokenInfo()
	access, err := generateToken(info, AccessToken)
	if err != nil {
		return "", "", err
	}
	refresh, err := generateToken(info, RefreshToken)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, Unauthorized("invalid token")
}

//goland:noinspection GoUnusedParameter
func keyFunc(token *jwt.Token) (i interface{}, err error) {
	return []byte(config.Config.SecretKey), nil
}

func GetClaims(c *fiber.Ctx) *MyClaims {
	return c.Locals("claims").(*MyClaims)
}
