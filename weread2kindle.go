package main

import (
	"flag"
	"log"
)

var (
	bookID int
	cookie string
	dir    string
)

func init() {
	flag.IntVar(&bookID, "b", 0, "book id, should be great than 0")
	flag.StringVar(&cookie, "c", "", "weread login cookie")
	flag.StringVar(&dir, "d", ".", "output directory")
	flag.Parse()
}

func main() {
	if bookID < 1 || cookie == "" {
		flag.PrintDefaults()
		return
	}
	book := NewBook(bookID)
	err := book.WereadBook(cookie)
	if err != nil {
		log.Fatal(err)
	}
	err = book.ToFiles(dir)
	if err != nil {
		log.Fatal(err)
	}
}
