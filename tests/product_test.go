package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	"testing"
)

func TestGetProduct(t *testing.T) {
	product := testAPIModel[Product](t, "get", "/api/products/1", 200)
	assert.Equal(t, 1, product.ID)
}

func TestListProduct(t *testing.T) {
	// return all products
	var length int64
	DB.Table("product").Where("tenant_id = ?", 1).Count(&length)
	resp := testAPIModel[[]Product](t, "get", "/api/categories/1/products", 200)
	assert.Equal(t, length, int64(len(resp)))
}

func TestAddProduct(t *testing.T) {
	data := Map{
		"name":        "TestAddProduct",
		"description": "TestAddProductDescription",
		"price":       100,
		"tenant_id":   1,
		"type":        -1,
	}
	testAPI(t, "post", "/api/categories/1/products", 201, data)

	// no price
	data = Map{
		"name":      "TestAddProduct",
		"tenant_id": 1,
		"type":      -1,
	}
	testAPI(t, "post", "/api/categories/1/products", 400, data)

	// no tenant_id
	data = Map{
		"name":  "TestAddProduct",
		"price": 100,
		"type":  -1,
	}
	testAPI(t, "post", "/api/categories/1/products", 400, data)

	// no type
	data = Map{
		"name":      "TestAddProduct",
		"price":     100,
		"tenant_id": 1,
	}
	testAPI(t, "post", "/api/categories/1/products", 400, data)
}

func TestModifyProduct(t *testing.T) {
	data := Map{
		"name":        "modify",
		"description": "modify",
		"price":       200,
		"closed":      true,
	}
	product := testAPIModel[Product](t, "put", "/api/products/1", 200, data)
	assert.Equal(t, "modify", product.Name)
	assert.Equal(t, "modify", product.Description)
	assert.Equal(t, 200, product.Price)
	assert.Equal(t, true, product.Closed)

	// modify closed to false
	data = Map{
		"closed": false,
	}
	product = testAPIModel[Product](t, "put", "/api/products/1", 200, data)
	assert.Equal(t, false, product.Closed)
	assert.Equal(t, "modify", product.Name)
	assert.Equal(t, "modify", product.Description)
	assert.Equal(t, 200, product.Price)
}

func TestDeleteProduct(t *testing.T) {
	testAPI(t, "delete", "/api/products/2", 204)
	var product Product
	DB.Find(&product, 2)
	assert.Equal(t, true, product.Closed)
}
