package model

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	Name string   `gorm:"uniqueIndex"`
	Uptime   []Uptime `gorm:"foreignkey:WebsiteID"`
}

type Uptime struct {
	gorm.Model
	WebsiteID uint
	Up        bool
	Ping      string
}
