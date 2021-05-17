package entity

import "gorm.io/gorm"

type LogEntity struct {
	gorm.Model
	Level   string
	Message string
}

func LogTable(log LogEntity) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(log.Level)
	}
}
