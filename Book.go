package main

import (
	"strconv"
)

type Bookface interface {
	PrintBook() string
}
type Book struct {
	title  string
	author string
	pages  int
}

func (b Book) PrintBook() string {

	result := "\n Title of the book : " + b.title + "\n Author of the book : " + b.author +
		"\n Pages : " + strconv.Itoa(b.pages)
	return result
}

func ConstructorBook(_title string, _author string, _pages int) *Book {
	b := new(Book)
	b.title = _title
	b.author = _author
	b.pages = _pages
	return b
}
