package product

import (
	"github.com/gofiber/fiber/v2"
	"keyi/auth"
	. "keyi/models"
	"keyi/utils"
)

// ListRecords
// @Summary List View Records of a User
// @Tags Record
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Router /products/records [get]
// @Success 200 {array} Product
func ListRecords(c *fiber.Ctx) error {
	var query Query
	err := utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	var products []Product
	DB.Limit(query.Size).Offset(query.Offset).Order("product_record.id "+query.Sort).
		Joins("INNER JOIN product_record ON product.id = product_record.product_id").
		Where("product_record.user_id = ?", auth.GetClaims(c).UID).
		Find(&products)

	return c.JSON(products)
}

// AddRecord
// @Summary Add a product to view record
// @Description Frontend should call this api each time user enters the detail page of a product.
// @Tags Record
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "product id"
// @Router /products/{id}/records [post]
// @Success 201 {object} ProductRecord
func AddRecord(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	record := ProductRecord{
		UserID:    auth.GetClaims(c).UID,
		ProductID: id,
	}

	err = DB.Create(&record).Error
	if err != nil {
		return err
	}

	return c.Status(201).JSON(record)
}

// DeleteRecord
// @Summary Delete a product view record
// @Tags Record
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "product id"
// @Router /products/{id}/records [delete]
// @Success 204
func DeleteRecord(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = DB.
		Where("user_id = ?", auth.GetClaims(c).UID).
		Where("product_id = ?", id).
		Delete(&ProductRecord{}).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}

// DeleteAllRecords
// @Summary Delete all product view records
// @Tags Record
// @Produce application/json
// @Security ApiKeyAuth
// @Router /products/records [delete]
// @Success 204
func DeleteAllRecords(c *fiber.Ctx) error {
	err := DB.
		Where("user_id = ?", auth.GetClaims(c).UID).
		Delete(&ProductRecord{}).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}
