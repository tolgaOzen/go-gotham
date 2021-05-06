package scopes

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gotham/helpers"
)

type Pagination struct {
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	OrderBy string `query:"order_by"`
	SortBy  string `query:"sort_by"`
}

func (r *Pagination) Paginate(tableName string, defaultField string, vars []interface{}, orderBy ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}

		if r.Limit <= 0 {
			r.Limit = 20
		}

		if !helpers.InArray(r.SortBy, []string{"asc", "desc"}) {
			r.SortBy = "asc"
		}

		if !helpers.InArray(r.OrderBy, orderBy) {
			return db.Clauses(clause.OrderBy{
				Expression: clause.Expr{SQL: fmt.Sprintf("FIELD(`%v`.`%v`,?) %v", tableName, defaultField, r.SortBy), Vars: vars, WithoutParentheses: true},
			}).Offset(helpers.OffsetCal(r.Page, r.Limit)).Limit(r.Limit)
		}

		if tableName == "" {
			return db.Order(fmt.Sprintf("`%v` %v", r.OrderBy, r.SortBy)).Offset(helpers.OffsetCal(r.Page, r.Limit)).Limit(r.Limit)
		} else {
			return db.Order(fmt.Sprintf("`%v`.`%v` %v", tableName, r.OrderBy, r.SortBy)).Offset(helpers.OffsetCal(r.Page, r.Limit)).Limit(r.Limit)
		}
	}
}
