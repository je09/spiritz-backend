package repository

import "github.com/je09/spritz-backend/internal/entity"

type Bookshelf interface {
	CreatePublic(user entity.User) error
	CreateLibrary() error
	Get(user entity.User) (entity.Bookshelf, error)
	GetAll() ([]entity.Bookshelf, error)
	Remove(vkID int64) error
}

type Book interface {
	Create(book entity.Book, shelf entity.Bookshelf) error
	Remove(book entity.Book) error
}

type Statistics interface {
}

type User interface {
	Create(entity.User) error
	CreateLibrarian() error
	Delete(entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(vkID int64) (*entity.User, error)
	Check(user entity.User) bool
}
