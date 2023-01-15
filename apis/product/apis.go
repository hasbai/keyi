package product

import (
	"github.com/gofiber/fiber/v2"
	"keyi/auth"
	. "keyi/models"
	"keyi/utils"
)

// GetProduct
// @Summary Get a product
// @Tags Product
// @Produce application/json
// @Param id path int true "id"
// @Router /products/{id} [get]
// @Success 200 {object} Product
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var product Product
	err = DB.Find(&product, id).Error
	if err != nil {
		return err
	}
	return c.JSON(product)
}

// ListProducts
// @Summary List Products
// @Description Return products that have the same talent_id as the user in a certain category
// @Tags Product
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Param category_id path int true "category_id"
// @Router /categories/{category_id}/products [get]
// @Success 200 {array} Product
func ListProducts(c *fiber.Ctx) error {
	var query Query
	err := utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	categoryID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var products []Product
	query.BaseQuery().
		Where("tenant_id = ?", auth.GetClaims(c).TenantID).
		Where("category_id = ?", categoryID).
		Where("closed = ?", false).
		Find(&products)

	return c.JSON(products)
}

// ListUserProducts
// @Summary List Products of a User
// @Tags Product
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Param id path int true "user id"
// @Router /users/{id}/products [get]
// @Success 200 {array} Product
func ListUserProducts(c *fiber.Ctx) error {
	var query ListUserProductsQuery
	err := utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	userID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = auth.OwnerOrPerm(c, auth.PAdmin, userID)
	if err != nil {
		return err
	}

	var products []Product
	clause := query.BaseQuery().Where("user_id = ?", userID)
	switch query.Closed {
	case 1:
		clause = clause.Where("closed = ?", false)
	case -1:
		clause = clause.Where("closed = ?", true)
	}
	if query.Type != ProductTypeAll {
		clause = clause.Where("type = ?", query.Type)
	}

	clause.Find(&products)

	return c.JSON(products)
}

// ListUserProductsType
// @Summary List products that a user bought or sold
// @Tags Product
// @Produce application/json
// @Security ApiKeyAuth
// @Param object query Query false "query"
// @Param id path int true "user id"
// @Param type path string true "bought or sold"
// @Router /users/{id}/products/{type} [get]
// @Success 200 {array} Product
func ListUserProductsType(c *fiber.Ctx) error {
	var query Query
	err := utils.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	userID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = auth.OwnerOrPerm(c, auth.PAdmin, userID)
	if err != nil {
		return err
	}

	var products []Product
	clause := query.BaseQuery().Where("closed = ?", true)

	switch c.Params("type") {
	case "sold":
		clause = clause.
			Where("type = ? AND user_id = ?", ProductTypeAsk, userID).
			Or("type = ? AND partner_id = ?", ProductTypeBid, userID)
	case "bought":
		clause = clause.
			Where("type = ? AND user_id = ?", ProductTypeBid, userID).
			Or("type = ? AND partner_id = ?", ProductTypeAsk, userID)
	default:
		return utils.BadRequest("invalid type, should be bought or sold")
	}

	clause.Find(&products)

	return c.JSON(products)
}

// AddProduct
// @Summary Add a product
// @Tags Product
// @Produce application/json
// @Security ApiKeyAuth
// @Param json body CreateModel true "json"
// @Param category_id path int true "category_id"
// @Router /categories/{category_id}/products [post]
// @Success 201 {object} Product
func AddProduct(c *fiber.Ctx) error {
	var body CreateModel
	err := utils.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	categoryID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	product := Product{
		Name:        body.Name,
		Description: body.Description,
		Images:      body.Images,
		Price:       body.Price,
		Type:        body.Type,
		CategoryID:  categoryID,
		UserID:      auth.GetClaims(c).UID,
	}

	err = DB.Create(&product).Error
	if err != nil {
		return err
	}

	return c.Status(201).JSON(product)
}

// ModifyProduct
// @Summary Modify a product
// @Tags Product
// @Produce application/json
// @Security ApiKeyAuth
// @Param json body ModifyModel true "json"
// @Param id path int true "product id"
// @Router /products/{id} [put]
// @Success 201 {object} Product
func ModifyProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var body ModifyModel
	err := utils.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	var product Product
	err = DB.
		Where("user_id = ?", auth.GetClaims(c).UID).
		Where("id = ?", id).
		First(&product).Error
	if err != nil {
		return err
	}

	bodyMap := utils.ToMap(&body)
	err = DB.Model(&product).Updates(&bodyMap).Error
	if err != nil {
		return err
	}

	return c.JSON(product)
}

// DeleteProduct
// @Summary Set a product as closed.
// @Tags Product
// @Produce application/json
// @Security ApiKeyAuth
// @Router /products/{id} [delete]
// @Param id path int true "product id"
// @Success 204
func DeleteProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var product Product
	err := DB.
		Where("user_id = ?", auth.GetClaims(c).UID).
		Where("id = ?", id).
		First(&product).Error
	if err != nil {
		return err
	}

	product.Closed = true
	err = DB.Save(&product).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}
