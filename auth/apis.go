package auth

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"io"
	"keyi/config"
	. "keyi/models"
	"keyi/utils"
	"net/http"
	"strings"
)

// Login
// @Summary Login
// @Description use wx code to get access token
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

	wxResponse, err := login(body.Code)
	if err != nil {
		return err
	}

	var user User
	user.OpenID = wxResponse.OpenID
	DB.FirstOrCreate(&user, "open_id = ?", wxResponse.OpenID)

	access, refresh, err := GenerateTokens(&user)
	if err != nil {
		return err
	}

	return c.JSON(TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

func login(code string) (WxResponse, error) {
	var wxResponse WxResponse

	if config.Config.Debug {
		wxResponse.OpenID = "test-openid"
		return wxResponse, nil
	}

	response, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.Config.AppID, config.Config.AppSecret, code,
	))
	if err != nil {
		return wxResponse, errors.New("wx login failed, http request failed")
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return wxResponse, errors.New("wx login failed, read response body failed")
	}

	err = json.Unmarshal(bodyBytes, &wxResponse)
	if err != nil {
		return wxResponse, errors.New("wx login failed, unmarshal response body failed")
	}
	if wxResponse.ErrorCode != 0 {
		return wxResponse, errors.New("wx login failed, " + wxResponse.ErrorMsg)
	}

	return wxResponse, nil
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
// @Description Fill in user's information to complete registration
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

	var user User
	err = DB.First(&user, c.Locals("claims").(MyClaims).UID).Error
	if err != nil {
		return err
	}

	if body.Username == "" {
		body.Username = utils.MD5(body.Email)
	}
	user.Email = body.Email
	user.Username = body.Username
	user.TenantID = body.TenantID
	user.TenantAreaID = body.TenantAreaID

	err = validateEmail(user)
	if err != nil {
		return err
	}

	return sendEmail(c, user)
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
	var tenant Tenant
	err := DB.First(&tenant, user.TenantID).Error
	if err != nil {
		return err
	}

	userDomain := strings.Split(user.Email, "@")[1]
	if !slices.Contains(tenant.Domains, userDomain) {
		return utils.BadRequest(
			"illegal email domain, please use one of the following: " +
				strings.Join(tenant.Domains, ", "),
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
	TenantID     int    `json:"tenant_id" validate:"required"`
	TenantAreaID int    `json:"tenant_area_id"`
}

type LoginBody struct {
	Code string `json:"code" validate:"required"`
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

type WxResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrorMsg   string `json:"errmsg"`
	ErrorCode  int    `json:"errcode"`
}
