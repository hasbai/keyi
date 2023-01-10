package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"keyi/config"
	. "keyi/models"
	"keyi/utils"
	"strings"
)

// Login
// @Summary Login
// @Description use email and password to get jwt tokens, valid user only
// @Tags Auth
// @Produce application/json
// @Param json body LoginBody true "json"
// @Router /login [post]
// @Success 200 {object} TokenResponse
func Login(c *fiber.Ctx) error {
	var body LoginBody
	err := utils.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	var user User
	if body.Email != "" {
		err = DB.Where("email = ?", body.Email).First(&user).Error
	} else {
		err = DB.Where("username = ?", body.Username).First(&user).Error
	}
	if err != nil {
		return utils.BadRequest("user does not exist")
	}

	if !user.IsValid {
		return utils.BadRequest("user is not activated, please check your email")
	}
	if user.Password != body.Password {
		return utils.BadRequest("password is incorrect")
	}

	access, refresh, err := GenerateTokens(&user)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

// Refresh
// @Summary Refresh Token
// @Description use refresh token to refresh tokens
// @Tags Auth
// @Produce application/json
// @Param json body RefreshBody true "json"
// @Router /refresh [post]
// @Success 200 {object} TokenResponse
func Refresh(c *fiber.Ctx) error {
	var body RefreshBody
	err := utils.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	claims, err := ParseToken(body.RefreshToken)
	if err != nil {
		return err
	}
	if claims.Type != TokenTypeRefresh || !claims.IsValid {
		return utils.BadRequest("invalid refresh token")
	}

	access, refresh, err := GenerateTokens(claims)
	if err != nil {
		return err
	}

	return c.JSON(TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

// Register
// @Summary Register
// @Tags Auth
// @Produce application/json
// @Param json body RegisterBody true "json"
// @Router /register [post]
// @Success 200 {object} MessageResponse
func Register(c *fiber.Ctx) error {
	var body RegisterBody
	err := utils.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	var emailExists bool
	DB.Raw("SELECT 1 FROM user WHERE email = ? LIMIT 1", body.Email).Scan(&emailExists)
	if emailExists {
		return utils.BadRequest("email already registered, please login")
	}

	user := User{
		Username:     body.Username,
		Email:        body.Email,
		Password:     body.Password,
		TEL:          body.TEL,
		TenantID:     body.TenantID,
		TenantAreaID: body.TenantAreaID,
	}
	if user.Username == "" {
		user.Username = utils.MD5(user.Email)
	}

	err = validateEmail(user)
	if err != nil {
		return err
	}

	err = DB.Create(&user).Error
	if err != nil {
		return err
	}

	err = sendEmail(c, user)
	if err != nil {
		return err
	}

	return nil
}

// Validate
// @Summary Validate By Email
// @Description send validation email to user
// @Tags Auth
// @Produce application/json
// @Param object query ValidateModel true "query"
// @Router /validate [post]
// @Success 200 {object} MessageResponse
func Validate(c *fiber.Ctx) error {
	var query ValidateModel
	err := utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	if query.Email != "" {
		var user User
		err = DB.Where("email = ?", query.Email).First(&user).Error
		if err != nil {
			return utils.BadRequest("user does not exist")
		}

		err = validateEmail(user)
		if err != nil {
			return err
		}

		return sendEmail(c, user)
	}

	// other validate methods

	return utils.BadRequest("email is required")
}

func validateEmail(user User) error {
	var domain string
	DB.Raw("SELECT domains FROM tenant WHERE id = ?", user.TenantID).Scan(&domain)
	if domain == "" {
		return utils.BadRequest("tenant not supported now")
	}
	domains := strings.Split(domain, ",")
	userDomain := strings.Split(user.Email, "@")[1]
	if !slices.Contains(domains, userDomain) {
		return utils.BadRequest(
			"illegal email domain, please use one of the following: " +
				strings.Join(domains, ", "),
		)
	}
	return nil
}

func sendEmail(c *fiber.Ctx, user User) error {
	code, err := setCode(user.ID)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"%s/api/users/%d/activate?code=%s",
		config.Config.BaseURL, user.ID, code,
	)
	text := fmt.Sprintf(`
		<h1>欢迎注册%s</h1>
		请点击链接激活账号：<br><br>
		<a href=%s>%s</a>`,
		config.Config.SiteName, url, url,
	)
	err = utils.SendEmail([]string{user.Email}, "注册验证", text)
	if err != nil {
		return err
	}

	return c.JSON(MessageResponse{
		"验证邮件已发送，请点击邮件中的链接完成注册，如果没有收到邮件，请检查垃圾箱",
	})
}

// Activate
// @Summary Activate
// @Description clicks the link in the email to activate the user
// @Tags Auth
// @Produce application/json
// @Param id path int true "user id"
// @Param object query ActivateQuery true "query"
// @Router /users/{id}/activate [get]
// @Success 200 {object} TokenResponse
func Activate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var query ActivateQuery
	err = utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	var user User
	err = DB.First(&user, id).Error
	if err != nil {
		return utils.BadRequest("user does not exist")
	}

	code, err := getCode(user.ID)
	if err != nil {
		return err
	}
	if query.Code != code {
		return utils.BadRequest("invalid code")
	}

	user.IsValid = true
	DB.Save(&user)

	access, refresh, err := GenerateTokens(&user)
	if err != nil {
		return err
	}

	return c.JSON(TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

type RegisterBody struct {
	Username     string `json:"username" validate:"max=32"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	TEL          string `json:"tel"`
	TenantID     int    `json:"tenant_id" validate:"required"`
	TenantAreaID int    `json:"tenant_area_id"`
}

type LoginBody struct {
	Username string `json:"username" validate:"max=32"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RefreshBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ActivateQuery struct {
	Code string `query:"code" validate:"required"`
}

type ValidateModel struct {
	Email string `query:"email" validate:"email"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
