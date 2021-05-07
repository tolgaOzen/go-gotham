package scopes

import (
	"gorm.io/gorm"
	"gotham/helpers"
	"gotham/utils"
)

type GormPager interface {
	ToPaginate() func(db *gorm.DB) *gorm.DB
}

type GormPagination struct {
	*utils.Pagination
}

func (r *GormPagination) ToPaginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(helpers.OffsetCal(r.Pagination.GetPage(), r.Pagination.GetLimit())).Limit(r.Pagination.GetLimit())
	}
}
