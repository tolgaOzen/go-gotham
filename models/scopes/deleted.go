package scopes

import 	"gorm.io/gorm"

func Deleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}
