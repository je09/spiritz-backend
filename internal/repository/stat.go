package repository

import (
	"gorm.io/gorm"
)

type StatisticsGorm struct {
	db *gorm.DB
}

func NewStatisticsGorm(db *gorm.DB) *StatisticsGorm {
	return &StatisticsGorm{db: db}
}
