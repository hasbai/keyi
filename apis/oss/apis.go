package oss

import (
	"github.com/gofiber/fiber/v2"
	"keyi/auth"
	"keyi/config"
)

func RegisterRoutes(app fiber.Router) {
	app.Get("/oss", GetOssInformation)
}

// GetOssInformation
// @Summary Get a category
// @Tags OSS
// @Produce application/json
// @Security ApiKeyAuth
// @Router /oss [get]
// @Success 200 {object} OssInformation
func GetOssInformation(c *fiber.Ctx) error {
	err := auth.PermOnly(c, auth.PUser)
	if err != nil {
		return err
	}

	return c.JSON(&OssInformation{
		Token:   genUploadToken(),
		Bucket:  config.Config.QiniuBucket,
		BaseURL: config.Config.QiniuBaseUrl,
		Expires: tokenExpires,
	})
}

type OssInformation struct {
	Token   string `json:"token"`
	Bucket  string `json:"bucket"`
	BaseURL string `json:"base_url"` // example: http://example.com/
	Expires int    `json:"expires"`  // in seconds
}
