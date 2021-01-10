package scopes

import (
	"gorm.io/gorm"
	"gotham/helpers"
	"gotham/requests"
)

func Paginate(r *requests.Pagination, model interface{}, orderDefault string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if r.Page == 0 {
			r.Page = 1
		}

		if r.Limit == 0 {
			r.Limit = 50
		}

		r.OrderBy = helpers.OrderBySetter(r.OrderBy, "query", model, orderDefault)

		if !helpers.InArray(r.SortBy, []string{"asc", "desc"}) {
			r.SortBy = "asc"
		}

		return db.Order(r.OrderBy + " " + r.SortBy).Offset(helpers.OffsetCal(r.Page, r.Limit)).Limit(r.Limit)
	}
}

