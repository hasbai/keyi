// Package models contains database models
package models

import (
	"database/sql/driver"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"keyi/config"
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

type StringSlice []string

func (t StringSlice) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *StringSlice) Scan(input any) error {
	return json.Unmarshal(input.([]byte), t)
}

// GormDataType gorm common data type
func (StringSlice) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
//goland:noinspection GoUnusedParameter
func (StringSlice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
