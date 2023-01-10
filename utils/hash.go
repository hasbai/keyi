package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	ret := md5.Sum([]byte(s))
	return hex.EncodeToString(ret[:])
}
