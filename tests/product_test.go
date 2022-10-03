package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	"strconv"
	"testing"
)

func TestGetProduct(t *testing.T) {
	product := testAPIModel[Product](t, "get", "/api/products/1", 200)
	assert.Equal(t, 1, product.ID)
}

func TestListProduct(t *testing.T) {
	var length int64
	DB.Table("product").Where("category_id = ?", 1).Count(&length)
	resp := testAPIArray(t, "get", "/api/categories/1/products", 200)
	assert.Equal(t, length, int64(len(resp)))
}

func TestAddProduct(t *testing.T) {
	data := Map{
		"name":        "TestAddProduct",
		"description": "TestAddProductDescription",
		"price":       1.0,
		"type":        1,
		"condition":   1.0,
		"location":    "TestAddProductLocation",
		"contact":     "TestAddProductContact",
	}
	resp := testAPIModel[Product](t, "post", "/api/categories/1/products", 201, data)
	assert.Equal(t, 1, resp.CategoryID)
}

func TestModifyProduct(t *testing.T) {
	data := Map{
		"name":        "TestModifyProduct",
		"description": "TestModifyProductDescription",
		"price":       1.0,
		"type":        1,
		"condition":   1.0,
		"location":    "TestModifyProductLocation",
		"contact":     "TestModifyProductContact",
	}
	resp := testAPIModel[Product](t, "put", "/api/products/1", 200, data)
	assert.Equal(t, 1, resp.CategoryID)
	assert.Equal(t, "TestModifyProduct", resp.Name)
	assert.Equal(t, "TestModifyProductDescription", resp.Description)
	assert.Equal(t, 1.0, resp.Price)
	assert.Equal(t, int8(1), resp.Type)
	assert.Equal(t, 1.0, resp.Condition)
	assert.Equal(t, "TestModifyProductLocation", resp.Location)
	assert.Equal(t, "TestModifyProductContact", resp.Contact)
}

func TestModifyProductClosed(t *testing.T) {
	product := Product{
		Name:   "TestModifyProductClosed",
		Closed: true,
	}
	DB.Create(&product)

	resp := testAPIModel[Product](
		t, "put",
		"/api/products/"+strconv.Itoa(product.ID),
		200,
		Map{"closed": false},
	)

	assert.Equal(t, false, resp.Closed)
}

func TestDeleteProduct(t *testing.T) {
	product := Product{
		Name: "TestDeleteProduct",
	}
	DB.Create(&product)
	testAPI(t, "delete", "/api/products/"+strconv.Itoa(product.ID), 204, nil)
	DB.First(&product, product.ID)
	assert.Equal(t, true, product.Closed)
}
