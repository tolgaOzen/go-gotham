package scopes

import (
	"fmt"

	"gorm.io/gorm"

	"gotham/helpers"
	"gotham/utils"
)

type GormOrderer interface {
	ToOrder(tableName string, defaultOrder string, orderByOptions ...string) func(db *gorm.DB) *gorm.DB
}

type GormOrder struct {
	*utils.Order
}

func (r *GormOrder) ToOrder(tableName string, defaultOrder string, orderByOptions ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if helpers.InArray(r.Order.GetOrderBy(), orderByOptions) {
			return db.Order(fmt.Sprintf("%v.%v %v", tableName, r.Order.GetOrderBy(), r.Order.GetSortBy()))
		}
		return db.Order(fmt.Sprintf("%v.%v asc", tableName, defaultOrder))
	}
}
