package bookparse

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	AvgWordLength   = 5
	AvgWordsPerPage = 1024
)

var (
	CharsPerPage = AvgWordsPerPage * AvgWordLength
)

type EBook struct {
	Book
}

func NewEBook(path string) (*EBook, error) {
	switch strings.ToLower(filepath.Ext(path)) {
	case ExtEPUB:
		return &EBook{Book: NewEPUBBook(path)}, nil
	default:
		return nil, fmt.Errorf("ebook format isn't supported")
	}
}

type Base struct {
	Author string
	Title  string

	Pages []string
}

// GetPage starts with 1.
func (b *Base) GetPage(page int) (string, int) {
	// Not in range.
	if page < 0 || page > len(b.Pages) {
		return "", 0
	}
	page--

	return b.Pages[page], b.countWords(page)
}

func (b *Base) IsEmpty() bool {
	return len(b.Pages) == 0
}

func (b *Base) TotalWords() int {
	words := 0
	for i := 0; i < len(b.Pages); i++ {
		words += b.countWords(i)
	}

	return words
}

func (b *Base) TotalPages() int {
	return len(b.Pages)
}

func (b *Base) countWords(page int) int {
	return strings.Count(b.Pages[page], " ")
}

func (b *Base) splitter(textBuffer string) (int, string) {
	if len(textBuffer) > CharsPerPage {
		// Index of the space after the max words.
		index := b.findEnd(textBuffer)
		// Put this text until the index in the paginator.
		b.Pages = append(b.Pages, textBuffer[:index])

		return b.splitter(textBuffer[index+1:])
	}

	return len(b.Pages), textBuffer
}

func (b *Base) findEnd(text string) int {
	for i, v := range text[CharsPerPage:] {
		if v == ' ' {
			return CharsPerPage + i
		}
	}

	return len(text) - 1
}
