package bookparse

import "io"

type Book interface {
	Parse() error
	GetPage(page int) (string, int)
	IsEmpty() bool
	TotalWords() int
	TotalPages() int
	countWords(page int) int
	parseMeta()
	parsePages() error
	parseText(r io.ReadCloser) (string, error)
	splitter(textBuffer string) (int, string)
	findEnd(text string) int
}
