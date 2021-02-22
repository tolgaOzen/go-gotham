package scopes

import (
	"gorm.io/gorm"
	"gotham/helpers"
)

type Pagination struct {
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	OrderBy string `query:"order_by"`
	SortBy  string `query:"sort_by"`
}

func (r *Pagination) Paginate(tableName string ,orderBy ...string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}

		if r.Limit <= 0 {
			r.Limit = 20
		}

		if !helpers.InArray(r.OrderBy, orderBy) {
			r.OrderBy = "id"
		}

		if !helpers.InArray(r.SortBy, []string{"asc", "desc"}) {
			r.SortBy = "asc"
		}

		return db.Order("`" + tableName + "`" + "." +"`" + r.OrderBy + "` " + r.SortBy).Offset(helpers.OffsetCal(r.Page, r.Limit)).Limit(r.Limit)
	}
}
