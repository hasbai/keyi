package models

import (
	"gorm.io/gorm"
	"keyi/config"
)

type CanQuery interface {
	BaseQuery() *gorm.DB
}

type Query struct {
	Size    int    `query:"size" default:"30" validate:"min=0,max=100"`    // length of object array
	Offset  int    `query:"offset" default:"0" validate:"min=0"`           // offset of object array
	Sort    string `query:"sort" default:"desc" validate:"oneof=asc desc"` // Sort order
	OrderBy string `query:"order_by" default:"id"`                         // SQL ORDER BY field
}

func (q *Query) BaseQuery() *gorm.DB {
	return config.DB.Limit(q.Size).Offset(q.Offset).Order(q.OrderBy + " " + q.Sort)
}

type MessageModel struct {
	Message string `json:"product"`
}
