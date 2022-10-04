package user

import (
	"fmt"
	"github.com/google/uuid"
	"keyi/utils"
	"math/rand"
	"time"
)

func randomCode() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000000)
	return fmt.Sprintf("%06d", random)
}

func defaultUsername() string {
	uid := uuid.NewString()
	return uid
}

func validateCode(email, code string) error {
	var myCode string
	err := utils.GetCache(email, &myCode)
	if err != nil {
		return err
	}
	if code == "" || myCode == "" {
		return utils.BadRequest("验证码不能为空")
	}
	if code != myCode {
		return utils.BadRequest("验证码错误")
	}
	return nil
}
