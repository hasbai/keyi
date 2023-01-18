package product

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	. "keyi/models"
	"regexp"
	"strings"
)

type ListProductsQuery struct {
	Query
	Search string `query:"search"`
	// 0: false, 1: true, -1: all, default is 0
	Closed int8 `query:"closed" validate:"min=-1,max=1"`
	// 0: all, 1: sell, -1: buy, default is 0
	Type int8 `query:"type" validate:"min=-1,max=1"`
}

func (q *ListProductsQuery) BaseQuery() *gorm.DB {
	query := DB

	search := q.Search != ""
	if search {
		q.OrderBy = "rank"
		query = query.Clauses(
			clause.Select{
				Columns: []clause.Column{
					{Name: "*", Raw: true},
					{Name: "ts_rank(tsv, query) AS rank", Raw: true},
				},
			},
			clause.From{
				Tables: []clause.Table{
					{Name: "product", Raw: true},
					{
						Name: fmt.Sprintf(
							"to_tsquery('chinese_zh', '%s') query",
							genSearch(q.Search),
						),
						Raw: true,
					},
				},
			},
		)
	}

	switch q.Closed {
	case 0:
		query = query.Where("closed = ?", false)
	case 1:
		query = query.Where("closed = ?", true)
	}

	if q.Type != ProductTypeAll {
		query = query.Where("type = ?", q.Type)
	}

	if search {
		query = query.Where("query @@ tsv")
	}

	return query.Limit(q.Size).Offset(q.Offset).Order(q.OrderBy + " " + q.Sort)
}

var pattern = regexp.MustCompile(`\s+`)

func genSearch(s string) string {
	return pattern.ReplaceAllString(strings.TrimSpace(s), " | ")
}
