// Package models contains database models
package models

import (
	"database/sql/driver"
	"keyi/config"
	"strings"
	"time"
)

var DB = config.DB

type Map = map[string]any

type BaseModel struct {
	ID        int       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model BaseModel) GetID() int {
	return model.ID
}

type StringArray []string // separate by space

func (t StringArray) Value() (driver.Value, error) {
	return strings.Join(t, " "), nil
}

func (t *StringArray) Scan(input any) error {
	str := input.(string)
	if str == "" {
		*t = []string{}
	} else {
		*t = strings.Split(str, " ")
	}
	return nil
}
