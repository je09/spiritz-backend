package repository

import "gorm.io/gorm"

type Repository struct {
	User
	Book
	Bookshelf
	Statistics
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:       NewUserGorm(db),
		Book:       NewBookGorm(db),
		Bookshelf:  NewBookshelfGorm(db),
		Statistics: NewStatisticsGorm(db),
	}
}
