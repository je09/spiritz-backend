package repository

import (
	"github.com/je09/spritz-backend/internal/entity"
	"github.com/je09/spritz-backend/internal/repository/models"
	"gorm.io/gorm"
)

type UserGorm struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) *UserGorm {
	return &UserGorm{db: db}
}

func (u *UserGorm) Create(user entity.User) error {
	r := u.db.Model(models.User{}).Create(user)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (u *UserGorm) CreateLibrarian() error {
	lbr := entity.User{VkID: 0}
	return u.Create(lbr)
}

func (u *UserGorm) Delete(user entity.User) error {
	r := u.db.Delete(models.User{}, user.VkID)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (u *UserGorm) GetAll() ([]entity.User, error) {
	var uu []entity.User
	r := u.db.Model(models.User{}).Find(&uu)
	if r.Error != nil {
		return nil, r.Error
	}

	return uu, nil
}

func (u *UserGorm) GetByID(vkID int64) (*entity.User, error) {
	var uu entity.User
	r := u.db.Model(models.User{}).First(&uu, vkID)
	if r.Error != nil {
		return nil, r.Error
	}

	return &uu, nil
}

func (u *UserGorm) Check(user entity.User) bool {
	var uu entity.User
	r := u.db.Model(models.User{}).Find(&uu, user.VkID)
	if r.Error != nil {
		return false
	}

	if uu.VkID == 0 {
		return false
	}

	return true
}
