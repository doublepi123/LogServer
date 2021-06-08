package entity

import "gorm.io/gorm"

type PictureEntity struct {
	gorm.Model
	OrgName  string
	Name     string
	FileName string
	Data     []byte
	Pid      string
}
