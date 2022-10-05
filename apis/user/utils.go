package user

import (
	"fmt"
	"github.com/google/uuid"
	"keyi/utils"
	"math/rand"
	"strings"
	"time"
)

func randomCode() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000000)
	return fmt.Sprintf("%06d", random)
}

func defaultUsername() string {
	uid := uuid.NewString()
	return strings.Replace(uid, "-", "", -1)
}

func validateCode(email, code string) error {
	if code == "" {
		return utils.BadRequest("验证码不能为空")
	}
	var myCode string
	err := utils.GetCache(email, &myCode)
	if err != nil {
		return err
	}
	if myCode == "" || code != myCode {
		return utils.BadRequest("验证码错误")
	}
	return nil
}
