package auth

import (
	"keyi/utils"
	"math/rand"
	"strconv"
	"time"
)

const codeExpires = time.Minute * 10

func codeCacheKey(id int) string {
	return "code-" + strconv.Itoa(id)
}

func setCode(id int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	code := strconv.FormatInt(rand.Int63(), 16)
	err := utils.SetCache(codeCacheKey(id), code, codeExpires)
	if err != nil {
		return "", err
	}
	return code, nil
}

func getCode(id int) (string, error) {
	var code string
	err := utils.GetCache(codeCacheKey(id), &code)
	if err != nil {
		return "", err
	}
	return code, nil
}
