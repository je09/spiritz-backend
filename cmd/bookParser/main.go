package main

import (
	"fmt"
	"github.com/je09/spritz-backend/pkg/bookparse"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	path := os.Args[1]
	book, err := bookparse.NewEBook(path)
	if err != nil {
		panic(err)
	}
	err = book.Parse()
	if err != nil {
		panic(err)
	}
	page, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Println(book.GetPage(page))
}
