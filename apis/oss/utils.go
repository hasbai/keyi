package oss

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"keyi/config"
)

const tokenExpires = 7200 // two hours

func genUploadToken() string {
	putPolicy := storage.PutPolicy{
		Scope: config.Config.QiniuBucket,
	}
	putPolicy.Expires = tokenExpires //示例2小时有效期
	mac := qbox.NewMac(config.Config.QiniuAccess, config.Config.QiniuSecret)
	return putPolicy.UploadToken(mac)
}
