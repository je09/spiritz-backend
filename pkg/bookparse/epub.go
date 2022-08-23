package bookparse

// Some parts of this package uses a code from @taylorskalyo goreader project.
// https://github.com/taylorskalyo/goreader.

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taylorskalyo/goreader/epub"
	"io"
	"regexp"
)

const (
	ExtEPUB = ".epub"
)

type EPUBBook struct {
	Base
	book     *epub.Rootfile
	filePath string
}

func NewEPUBBook(filePath string) *EPUBBook {
	return &EPUBBook{filePath: filePath}
}

func (b *EPUBBook) Parse() error {
	rc, err := epub.OpenReader(b.filePath)
	if err != nil {
		return err
	}
	defer rc.Close()

	// The rootfile (content.opf) lists all the contents of an epub file.
	// There may be multiple rootfiles, although typically there is only one.
	b.book = rc.Rootfiles[0]
	err = b.parsePages()
	if err != nil {
		return err
	}
	b.parseMeta()

	return nil
}

func (b *EPUBBook) parseMeta() {
	b.Author = b.book.Creator
	b.Title = b.book.Title
}

func (b *EPUBBook) parsePages() error {
	textBuffer := ""

	multipleSpacesRegexp := regexp.MustCompile("((\\n|\\s)+)")
	escapeSeqRegexp := regexp.MustCompile("(\\r|\\t|\\v)")

	for ch := 0; ch < len(b.book.Itemrefs); ch++ {
		f, err := b.book.Itemrefs[ch].Open()
		if err != nil {
			return err
		}

		textBuffer, err = b.parseText(f)
		if err != nil {
			return err
		}

		// Not very a efficient way of clearing strings,
		// But I just don't have any other idea.
		textBuffer = multipleSpacesRegexp.ReplaceAllString(textBuffer, " ")
		textBuffer = escapeSeqRegexp.ReplaceAllString(textBuffer, "")

		_, textBuffer = b.splitter(textBuffer)
	}
	b.Pages = append(b.Pages, textBuffer)

	return nil
}

func (b *EPUBBook) parseText(r io.ReadCloser) (string, error) {
	defer r.Close()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	return doc.Text(), nil
}
