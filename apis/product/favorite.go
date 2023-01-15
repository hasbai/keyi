package product

import (
	"github.com/gofiber/fiber/v2"
	"keyi/auth"
	. "keyi/models"
	"keyi/utils"
)

// ListFavorites
// @Summary List Favored Products of a User
// @Tags Favorite
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Router /products/favorite [get]
// @Success 200 {array} Product
func ListFavorites(c *fiber.Ctx) error {
	var query Query
	err := utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	var products []Product
	DB.Limit(query.Size).Offset(query.Offset).Order("product_favorite.id "+query.Sort).
		Joins("INNER JOIN product_favorite ON product.id = product_favorite.product_id").
		Where("product_favorite.user_id = ?", auth.GetClaims(c).UID).
		Find(&products)

	return c.JSON(products)
}

// AddFavorite
// @Summary Add a favored product
// @Tags Favorite
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "product id"
// @Router /products/{id}/favorite [post]
// @Success 201
func AddFavorite(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = DB.Create(&ProductFavorite{
		UserID:    auth.GetClaims(c).UID,
		ProductID: id,
	}).Error
	if err != nil {
		return err
	}

	return c.Status(201).JSON(nil)
}

// DeleteFavorite
// @Summary Delete a favored product
// @Tags Favorite
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "product id"
// @Router /products/{id}/favorite [delete]
// @Success 204
func DeleteFavorite(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = DB.
		Where("user_id = ?", auth.GetClaims(c).UID).
		Where("product_id = ?", id).
		Delete(&ProductFavorite{}).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}

// DeleteAllFavorites
// @Summary Delete all favored products
// @Tags Favorite
// @Produce application/json
// @Security ApiKeyAuth
// @Router /products/favorite [delete]
// @Success 204
func DeleteAllFavorites(c *fiber.Ctx) error {
	err := DB.
		Where("user_id = ?", auth.GetClaims(c).UID).
		Delete(&ProductFavorite{}).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}
