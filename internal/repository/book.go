package repository

import (
	"github.com/je09/spritz-backend/internal/entity"
	"github.com/je09/spritz-backend/internal/repository/models"
	"gorm.io/gorm"
)

type BookGorm struct {
	db *gorm.DB
}

func NewBookGorm(db *gorm.DB) *BookGorm {
	return &BookGorm{db: db}
}

func (b *BookGorm) Create(book entity.Book, shelf entity.Bookshelf) error {
	r := b.db.Model(models.Book{}).Create(book)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (b *BookGorm) Remove(book entity.Book) error {
	r := b.db.Model(models.Book{}).Delete(book)
	if r.Error != nil {
		return r.Error
	}

	return nil
}
