package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

//ToFiles out opf, html, ncx files
func (book *Book) ToFiles(dir string) error {
	opf, err := os.OpenFile(dir+"/book.opf", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	ncx, err := os.OpenFile(dir+"/book.ncx", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	_book, err := os.OpenFile(dir+"/book.html", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	err = opfTmp.Execute(opf, map[string]interface{}{
		"Book":      book,
		"XmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
	})
	if err != nil {
		return err
	}
	err = ncxTmp.Execute(ncx, map[string]interface{}{
		"Book":      book,
		"XmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
	})
	if err != nil {
		return err
	}
	err = bookTmp.Execute(_book, map[string]interface{}{
		"Book": book,
	})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dir+"/cover"+book.CoverImgExt, book.CoverImgData, 0600)
}
