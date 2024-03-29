package models

import (
	"fmt"
)

// models must be registered here to migrate into the database
func init() {
	fmt.Println("migrate database...")
	err := DB.AutoMigrate(
		&Category{},
		&Product{},
		&ProductRecord{},
		&ProductFavorite{},
	)
	if err != nil {
		panic(err)
	}
}
