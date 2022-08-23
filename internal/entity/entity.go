package entity

import (
	"github.com/gofrs/uuid"
	"os"
)

type User struct {
	VkID int64 `json:"vkID"`
}

type UserInfo struct {
	User  *User
	Books []Book
}

type Bookshelf struct {
	UUID   uuid.UUID `json:"UUID"`
	Public bool      `json:"Public"`
	User   *User     `json:"User"`
}

type Book struct {
	UUID uuid.UUID `json:"UUID"`

	Shelf *Bookshelf `gorm:"foreignKey:Bookshelf"`

	Title    string
	Author   string
	Pages    int64
	Words    int64
	Category *Category
	Public   bool

	FilePath os.File
}

type BookText struct {
	UUID uuid.UUID `json:"UUID"`

	Shelf *Bookshelf `gorm:"foreignKey:Bookshelf"`

	Title    string
	Author   string
	Pages    int64
	Words    int64
	Category *Category
	Public   bool

	Text string
}

type BookPage struct {
	Text  string
	Page  int
	Pages int
	Words int
}

type Category struct {
	Name  string
	Title string
}
