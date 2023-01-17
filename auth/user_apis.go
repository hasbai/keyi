package auth

import (
	"github.com/gofiber/fiber/v2"
	. "keyi/models"
	"keyi/utils"
)

// GetUser
// @Summary Get a user
// @Tags User
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "user id"
// @Router /users/{id} [get]
// @Success 200 {object} User
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = OwnerOrPerm(c, PAdmin, id)
	if err != nil {
		return err
	}

	var user User
	err = DB.First(&user, id).Error
	if err != nil {
		return err
	}

	return c.JSON(&user)
}

// ListUsers
// @Summary List users
// @Tags User
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Router /users [get]
// @Success 200 {array} User
func ListUsers(c *fiber.Ctx) error {
	err := PermOnly(c, PAdmin)
	if err != nil {
		return err
	}

	var query Query
	err = utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	var users []User
	err = query.BaseQuery().Find(&users).Error
	if err != nil {
		return err
	}

	return c.JSON(users)
}

// ModifyUser
// @Summary Modify a user
// @Tags User
// @Produce application/json
// @Security ApiKeyAuth
// @Param json body ModifyUserModel true "json"
// @Param id path int true "user id"
// @Router /users/{id} [put]
// @Success 201 {object} User
func ModifyUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = OwnerOrPerm(c, PAdmin, id)
	if err != nil {
		return err
	}

	var body ModifyUserModel
	err = utils.ValidateBody(c, &body)
	if err != nil {
		return err
	}
	// TODO: username modify limit

	var user User
	err = DB.First(&user, id).Error
	if err != nil {
		return err
	}

	err = DB.Model(&user).Updates(&body).Error
	if err != nil {
		return err
	}

	return c.JSON(&user)
}

// ListFollow
// @Summary List a user's follow users
// @Tags Follow
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Param id path int true "user id"
// @Router /users/{id}/follow [get]
// @Success 200 {array} User
func ListFollow(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = OwnerOrPerm(c, PAdmin, id)
	if err != nil {
		return err
	}

	var query Query
	err = utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	var users []User
	err = query.BaseQuery().
		Model(&User{BaseModel: BaseModel{ID: id}}).
		Association("Follow").
		Find(&users)
	if err != nil {
		return err
	}

	return c.JSON(users)
}

// ListFollowedBy
// @Summary List users follow a user
// @Tags Follow
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Param id path int true "user id"
// @Router /users/{id}/followed-by [get]
// @Success 200 {array} User
func ListFollowedBy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = OwnerOrPerm(c, PAdmin, id)
	if err != nil {
		return err
	}

	var query Query
	err = utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	var users []User
	query.BaseQuery().Joins(`
		INNER JOIN user_follow ON user_follow.user_id = "user".id
		AND user_follow.follow_id = ?`, id,
	).Find(&users)
	if err != nil {
		return err
	}

	return c.JSON(users)
}

// AddFollow
// @Summary Add a follow user
// @Tags Follow
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "user id"
// @Param f_id path int true "follow user id"
// @Router /users/{id}/follow/{f_id} [post]
// @Success 201
func AddFollow(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	followID, err := c.ParamsInt("f_id")
	if err != nil {
		return err
	}

	err = OwnerOrPerm(c, PAdmin, id)
	if err != nil {
		return err
	}

	err = DB.
		Model(&User{BaseModel: BaseModel{ID: id}}).
		Association("Follow").
		Append(&User{BaseModel: BaseModel{ID: followID}})
	if err != nil {
		return err
	}

	return c.Status(201).JSON(nil)
}

// DeleteFollow
// @Summary Delete a follow user
// @Tags Follow
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "user id"
// @Param f_id path int true "follow user id"
// @Router /users/{id}/follow/{f_id} [delete]
// @Success 204
func DeleteFollow(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	followID, err := c.ParamsInt("f_id")
	if err != nil {
		return err
	}

	err = OwnerOrPerm(c, PAdmin, id)
	if err != nil {
		return err
	}

	err = DB.
		Model(&User{BaseModel: BaseModel{ID: id}}).
		Association("Follow").
		Delete(&User{BaseModel: BaseModel{ID: followID}})
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}

type ModifyUserModel struct {
	Username     string `json:"username" validate:"max=32"`
	TenantID     int    `json:"tenant_id"`
	TenantAreaID int    `json:"tenant_area_id"`
	Description  string `json:"description"`
	Avatar       string `json:"avatar"`
	Contacts     string `json:"contacts"`
}
