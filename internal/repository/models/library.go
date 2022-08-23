package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"os"
)

type Bookshelf struct {
	gorm.Model

	UUID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Public bool
	User   User `gorm:"foreignKey:User"`
}

type Book struct {
	gorm.Model

	UUID  uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Shelf *Bookshelf `gorm:"foreignKey:Bookshelf"`

	Title    string
	Author   string
	Pages    int64
	Words    int64
	Category *Category
	Public   bool

	FilePath os.File
}

type Category struct {
	gorm.Model

	Name  string
	Title string
}

type Statistics struct {
	gorm.Model

	Book  *Book
	Shelf *Bookshelf

	Percentage float32
	PagesRead  int64
	WordsRead  int64
	AvrSpeed   float32
}
