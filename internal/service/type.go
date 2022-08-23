package service

import "github.com/je09/spritz-backend/internal/entity"

type User interface {
	Info(user entity.User) (entity.UserInfo, error)
	Create(user entity.User) error
	Remove(user entity.User) error
	GetAll(entity.User, error)
	Check(user entity.User) bool
}

type Category interface {
	GetAll() ([]entity.Category, error)
	Info() (entity.Category, error)
	GetBooks(category entity.Category) ([]entity.Book, error)
}

type Book interface {
	Page(book entity.Book, page int) (entity.BookPage, error)
}

type Library interface {
	Category
	Book

	Get(user entity.User) ([]entity.Book, error)
	GetPublic() ([]entity.Book, error)
	Create(user entity.User, book entity.Book) error
	CreateText(user entity.User, text entity.BookText) error
	CreatePublic(book entity.Book) error
	Borrow(book entity.Book) error
	Remove(user entity.User, book entity.Book) error
}
