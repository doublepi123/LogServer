package entity

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	Name string
}

type LogReturn struct {
	CreatedAt time.Time
	Level     string
	Message   string
}

type ItemReturn struct {
	Name string
}
type LogEntity struct {
	gorm.Model
	Level   string
	Message string
}

type LogRequest struct {
	Level string
	From  int
	To    int
}

func LogTable(log LogEntity) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(log.Level)
	}
}
