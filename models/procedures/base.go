package procedures

import (
	"gorm.io/gorm"
	"gotham/app"
)

type ProcedureI interface {
	drop(db *gorm.DB) error
	dropIfExist(db *gorm.DB) error
	create(db *gorm.DB) error
}

func CreateProcedure(p ProcedureI, db *gorm.DB) error {
	return p.create(db)
}

func DropProcedure(p ProcedureI, db *gorm.DB) error {
	return p.drop(db)
}

func DropProcedureIfExist(p ProcedureI, db *gorm.DB) error {
	return p.dropIfExist(db)
}

func Initialize() {
	db := app.Application.Container.GetDb().DB()
	_ = DropProcedureIfExist(UserCount{}, db)
	_ = CreateProcedure(UserCount{}, db)
}
