package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	"strconv"
	"testing"
)

func init() {
	categories := make([]Category, 10)
	for i := 1; i <= 10; i++ {
		categories[i-1].ID = i
		categories[i-1].Name = strconv.Itoa(i)
		categories[i-1].Description = strconv.Itoa(i)
	}
	DB.Create(&categories)

	products := make([]Product, 10)
	for i := 0; i < 10; i++ {
		products[i] = Product{
			Name:       "product-" + strconv.Itoa(i),
			CategoryID: 1,
			UserID:     1,
			TenantID:   1,
		}
	}
	DB.Create(&products)

	records := make([]ProductRecord, 0, 6)
	for _, uid := range []int{1, 2} {
		for _, id := range []int{1, 3, 5} {
			records = append(records, ProductRecord{
				UserID:    uid,
				ProductID: id,
			})
		}
	}
	DB.Create(&records)

	favorites := make([]ProductFavorite, 0, 6)
	for _, uid := range []int{1, 2} {
		for _, id := range []int{1, 3, 5} {
			favorites = append(favorites, ProductFavorite{
				UserID:    uid,
				ProductID: id,
			})
		}
	}
	DB.Create(&favorites)
}

func TestIndex(t *testing.T) {
	testCommon(t, "get", "/", 302)
	testCommon(t, "get", "/api", 200)
	data := testAPI(t, "get", "/404", 404)
	assert.Equal(t, "Cannot GET /404", data["message"])
}

func TestDocs(t *testing.T) {
	testCommon(t, "get", "/docs", 302)
	testCommon(t, "get", "/docs/index.html", 200)
}
