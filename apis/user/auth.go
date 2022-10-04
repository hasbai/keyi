package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"keyi/config"
	. "keyi/models"
	. "keyi/utils"
	"strings"
	"time"
)

const VerifyCodeExpire = time.Minute * 10

// VerifyEmail
// @Summary Verify a user's email
// @Tags Auth
// @Produce application/json
// @Router /verify/email [post]
// @Param json body EmailModel true "json"
// @Success 200 {object} MessageModel
func VerifyEmail(c *fiber.Ctx) error {
	var body EmailModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	emailWhiteList := []string{"fudan.edu.cn", "m.fudan.edu.cn"}
	domain := strings.Split(body.Email, "@")[1]
	if !slices.Contains(emailWhiteList, domain) {
		return BadRequest("Email domain not allowed")
	}

	code := randomCode()
	err = SetCache(body.Email, code, VerifyCodeExpire)
	if err != nil {
		return err
	}

	content := fmt.Sprintf("欢迎注册%s，您的验证码为：%s，有效期为10分钟。", config.Config.SiteName, code)
	err = SendEmail(fmt.Sprintf("%s注册验证", config.Config.SiteName), content, []string{body.Email})
	if err != nil {
		return err
	}

	message := fmt.Sprintf("验证码已发送至您的邮箱%s，请注意查收，如不存在，请检查垃圾邮件或稍等片刻", body.Email)
	return c.JSON(MessageModel{Message: message})
}
