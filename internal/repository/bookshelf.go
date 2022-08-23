package repository

import (
	"github.com/je09/spritz-backend/internal/entity"
	"github.com/je09/spritz-backend/internal/repository/models"
	"gorm.io/gorm"
)

type BookshelfGorm struct {
	db *gorm.DB
}

func NewBookshelfGorm(db *gorm.DB) *BookshelfGorm {
	return &BookshelfGorm{db: db}
}

func (b *BookshelfGorm) create(user entity.User, public bool) error {
	var uu models.User
	r := b.db.Find(&uu, user.VkID)
	if r.Error != nil {
		return r.Error
	}

	r = b.db.Create(models.Bookshelf{Public: public, User: uu})
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (b *BookshelfGorm) CreatePublic(user entity.User) error {
	return b.create(user, false)
}

func (b *BookshelfGorm) CreateLibrary() error {
	lbr := entity.User{VkID: 0}
	return b.create(lbr, true)
}

func (b *BookshelfGorm) Get(user entity.User) (entity.Bookshelf, error) {
	var bb entity.Bookshelf
	r := b.db.First(&bb, user.VkID)
	if r.Error != nil {
		return entity.Bookshelf{}, r.Error
	}

	return bb, nil
}

func (b *BookshelfGorm) GetAll() ([]entity.Bookshelf, error) {
	var bb []entity.Bookshelf
	r := b.db.Find(&bb)
	if r.Error != nil {
		return nil, r.Error
	}

	return bb, nil
}

func (b *BookshelfGorm) Remove(vkID int64) error {
	r := b.db.Delete(models.Bookshelf{}, vkID)
	if r.Error != nil {
		return r.Error
	}

	return nil
}
