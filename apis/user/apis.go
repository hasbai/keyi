package user

import (
	"github.com/gofiber/fiber/v2"
	. "keyi/models"
	. "keyi/utils"
)

// RegisterUser
// @Summary Register a new user
// @Tags User
// @Produce application/json
// @Param json body RegisterModel true "json"
// @Router /users [post]
// @Success 201 {object} User
func RegisterUser(c *fiber.Ctx) error {
	var body RegisterModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	err = validateCode(body.Email, body.Code)
	if err != nil {
		return err
	}

	var user User
	result := DB.Where("email = ?", body.Email).First(&user)
	if result.RowsAffected != 0 {
		return BadRequest("该邮箱已注册，请登录")
	}

	err = Copy(&body, &user, "Code")
	if err != nil {
		return err
	}
	if user.Name == "" {
		user.Name = defaultUsername()
	}

	err = DB.Create(&user).Error
	if err != nil {
		return err
	}

	return c.Status(201).JSON(&user)
}

// ModifyUser
// @Summary Modify a user
// @Tags User
// @Produce application/json
// @Router /users/{id} [put]
// @Param id path int true "id"
// @Param json body ModifyModel true "json"
// @Success 200 {object} User
func ModifyUser(c *fiber.Ctx) error {
	var body ModifyModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	// TODO: owner or admin

	var user User
	err = DB.Where("id = ?", c.Params("id")).First(&user).Error
	if err != nil {
		return err
	}

	if body.NewPassword != "" {
		if body.Password != "" { // password validation
			if !user.CheckPassword(body.Password) {
				return BadRequest("密码错误")
			}
		} else { // email validation
			err = validateCode(user.Email, body.Code)
			if err != nil {
				return err
			}
		}
		user.Password = body.NewPassword
	}

	err = Copy(&body, &user, "Code", "Password", "NewPassword")
	if err != nil {
		return err
	}

	err = DB.Save(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(&user)
}
