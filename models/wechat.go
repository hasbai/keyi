package models

import (
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/auth"
	"keyi/config"
)

var WeSDK *auth.Auth

func init() {
	sdk := weapp.NewClient(config.Config.AppID, config.Config.AppSecret)
	WeSDK = sdk.NewAuth()
}
