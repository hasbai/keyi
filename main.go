package main

import (
	"keyi/app"
	"keyi/utils"
)

// @title Ke yi
// @version 0.1.0
// @description This is a campus flea market system.

// @contact.name Maintainer Shi Yue
// @contact.email jsclndnz@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	myApp := app.Create()
	//goland:noinspection GoUnhandledErrorResult
	defer utils.Logger.Sync()
	err := myApp.Listen("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
