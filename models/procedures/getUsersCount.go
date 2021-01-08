package procedures

import (
	"gorm.io/gorm"
)

type UserCount struct {
	Count int  `json:"rate"`
}

func (UserCount) create(db *gorm.DB) error {
	sql := `CREATE PROCEDURE GetUsersCount()
	BEGIN
	 SELECT COUNT(*) as count FROM users;
	END`
	return db.Exec(sql).Error
}

func (UserCount) drop(db *gorm.DB) error {
	sql := `DROP PROCEDURE GetUserCount;`
	return db.Exec(sql).Error
}

func (UserCount) dropIfExist(db *gorm.DB) error {
	sql := `DROP PROCEDURE IF EXISTS GetUserCount;`
	return db.Exec(sql).Error
}

func GetUserCount(db *gorm.DB) UserCount {
	var returnVal UserCount
	db.Raw("CALL GetUserCount()").Scan(&returnVal)
	return returnVal
}

