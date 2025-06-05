package entity

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDb(gormDb *gorm.DB) {
	db = gormDb
}

func Db() *gorm.DB {
	if db == nil {
		log.Panic("entity.Db: database is not set")
	}

	return db
}
