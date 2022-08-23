package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	VkID      int64     `gorm:"primaryKey"`
	CreatedAt time.Time ``
}
