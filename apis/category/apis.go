package category

import (
	"github.com/gofiber/fiber/v2"
	. "keyi/models"
	. "keyi/utils"
)

// ListCategories
// @Summary List Categories
// @Tags Category
// @Produce application/json
// @Router /categories [get]
// @Success 200 {array} Category
func ListCategories(c *fiber.Ctx) error {
	var categories []Category
	DB.Find(&categories)
	return c.JSON(categories)
}

// GetCategory
// @Summary Get a category
// @Tags Category
// @Produce application/json
// @Param id path int true "id"
// @Router /categories/{id} [get]
// @Success 200 {object} Category
func GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var category Category
	err = DB.Find(&category, id).Error
	if err != nil {
		return err
	}
	return c.JSON(category)
}

// AddCategory
// @Summary Add A Category
// @Tags Category
// @Produce application/json
// @Param json body CreateModel true "json"
// @Router /categories [post]
// @Success 200 {object} Category
func AddCategory(c *fiber.Ctx) error {
	var body CreateModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	category := Category{
		Name:        body.Name,
		Description: body.Description,
	}
	result := DB.FirstOrCreate(
		&category,
		Category{Name: category.Name},
	)
	if result.RowsAffected == 0 {
		c.Status(200)
	} else {
		c.Status(201)
	}

	return c.JSON(category)
}

// ModifyCategory
// @Summary Modify A Category
// @Tags Category
// @Produce application/json
// @Param id path int true "id"
// @Param json body CreateModel true "json"
// @Router /categories/{id} [put]
// @Success 200 {object} Category
func ModifyCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var body CreateModel
	err = ValidateBody(c, &body)
	if err != nil {
		return err
	}

	var category Category
	err = DB.First(&category, id).Error
	if err != nil {
		return err
	}

	err = DB.Model(&category).Updates(&Category{
		Name:        body.Name,
		Description: body.Description,
	}).Error
	if err != nil {
		return err
	}

	return c.JSON(category)
}

// DeleteCategory
// @Summary Delete a category
// @Tags Category
// @Produce application/json
// @Param json body DeleteModel true "json"
// @Param id path int true "id"
// @Router /categories/{id} [delete]
// @Success 204
func DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var body DeleteModel
	err = ValidateBody(c, &body)
	if err != nil {
		return err
	}

	if id == body.To {
		return BadRequest("The deleted category can't be the same as to.")
	}

	err = DB.Exec("UPDATE product SET category_id = ? WHERE category_id = ?", body.To, id).Error
	if err != nil {
		return err
	}
	err = DB.Delete(&Category{}, id).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}
