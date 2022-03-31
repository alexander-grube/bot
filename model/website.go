package model

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Up   bool
}
