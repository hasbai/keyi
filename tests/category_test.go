package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	"strconv"
	"testing"
)

func init() {
	DB.Create(&User{
		BaseModel: BaseModel{
			ID: 1,
		}})
	categories := make([]Category, 10)
	for i := 1; i <= 10; i++ {
		categories[i-1].ID = i
		categories[i-1].Name = strconv.Itoa(i)
		categories[i-1].Description = strconv.Itoa(i)
	}
	products := make([]Product, 10)
	for i := 0; i < 10; i++ {
		products[i] = Product{
			CategoryID: 1,
			UserID:     1,
		}
	}
	DB.Create(&categories)
	DB.Create(&products)
}

func TestGetCategory(t *testing.T) {
	category := testAPIModel[Category](t, "get", "/api/categories/1", 200)
	assert.Equal(t, 1, category.ID)
}

func TestListCategory(t *testing.T) {
	// return all categories
	var length int64
	DB.Table("category").Count(&length)
	resp := testAPIArray(t, "get", "/api/categories", 200)
	assert.Equal(t, length, int64(len(resp)))
}

func TestAddCategory(t *testing.T) {
	data := Map{"name": "TestAddCategory", "description": "TestAddCategoryDescription"}
	testAPI(t, "post", "/api/categories", 201, data)

	// duplicate post, return 200 and change nothing
	data["description"] = "another"
	resp := testAPI(t, "post", "/api/categories", 200, data)
	assert.Equal(t, "TestAddCategoryDescription", resp["description"])
}

func TestModifyCategory(t *testing.T) {
	pinned := []int{3, 2, 5, 1, 4}
	data := Map{"name": "modify", "description": "modify", "pinned": pinned}

	category := testAPIModel[Category](t, "put", "/api/categories/1", 200, data)

	// test modify
	assert.Equal(t, "modify", category.Name)
	assert.Equal(t, "modify", category.Description)
}

func TestDeleteCategory(t *testing.T) {
	id := 3
	toID := 2

	hole := Product{CategoryID: id}
	DB.Create(&hole)
	testAPI(t, "delete", "/api/categories/"+strconv.Itoa(id), 204, Map{"to": toID})
	testAPI(t, "delete", "/api/categories/"+strconv.Itoa(id), 204, Map{"to": toID}) // repeat delete

	// deleted
	var d Category
	result := DB.First(&d, id)
	assert.True(t, result.Error != nil)

	// hole moved
	DB.First(&hole, hole.ID)
	assert.Equal(t, toID, hole.CategoryID)

}
