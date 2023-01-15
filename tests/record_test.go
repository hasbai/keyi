package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	"testing"
)

func TestListRecords(t *testing.T) {
	products := testAPIModel[[]Product](t, "get", "/api/products/records", 200)
	if len(products) != 3 {
		t.Error("Expected 3 records, got", len(products))
	}
	ids := make([]int, 3)
	for i, product := range products {
		ids[i] = product.ID
	}
	assert.Equal(t, []int{5, 3, 1}, ids)
}

func TestAddRecord(t *testing.T) {
	testAPI(t, "post", "/api/products/2/records", 201)
	var record ProductRecord
	err := DB.Where("user_id = ? AND product_id = ?", 1, 2).First(&record).Error
	if err != nil {
		t.Error("Failed to create record", err)
	}
}

func TestDeleteRecord(t *testing.T) {
	record := ProductRecord{
		UserID:    1,
		ProductID: 9,
	}
	DB.Create(&record)

	testAPI(t, "delete", "/api/products/1/records", 204)

	var result ProductRecord
	err := DB.Where("user_id = ? AND product_id = ?", 1, 1).First(&result).Error
	if err == nil {
		t.Error("Failed to delete record")
	}
}

func TestDeleteAll(t *testing.T) {
	setToken(2)
	testAPI(t, "delete", "/api/products/records", 204)
	resetToken()

	var records []ProductRecord
	DB.Where("user_id = ?", 2).Find(&records)
	if len(records) != 0 {
		t.Error("Failed to delete all records")
	}
}
