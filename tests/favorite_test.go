package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	"testing"
)

func TestListFavorites(t *testing.T) {
	products := testAPIModel[[]Product](t, "get", "/api/products/favorite", 200)
	if len(products) != 3 {
		t.Error("Expected 3 favorites, got", len(products))
	}
	ids := make([]int, 3)
	for i, product := range products {
		ids[i] = product.ID
	}
	assert.Equal(t, []int{5, 3, 1}, ids)
}

func TestAddFavorite(t *testing.T) {
	testAPI(t, "post", "/api/products/2/favorite", 201)
	var favorite ProductFavorite
	err := DB.Where("user_id = ? AND product_id = ?", 1, 2).First(&favorite).Error
	if err != nil {
		t.Error("Failed to create favorite", err)
	}
}

func TestDeleteFavorite(t *testing.T) {
	favorite := ProductFavorite{
		UserID:    1,
		ProductID: 9,
	}
	DB.Create(&favorite)

	testAPI(t, "delete", "/api/products/1/favorite", 204)

	var result ProductFavorite
	err := DB.Where("user_id = ? AND product_id = ?", 1, 1).First(&result).Error
	if err == nil {
		t.Error("Failed to delete favorite")
	}
}

func TestDeleteAllFavorites(t *testing.T) {
	setToken(2)
	testAPI(t, "delete", "/api/products/favorite", 204)
	resetToken()

	var favorites []ProductFavorite
	DB.Where("user_id = ?", 2).Find(&favorites)
	if len(favorites) != 0 {
		t.Error("Failed to delete all favorites")
	}
}
