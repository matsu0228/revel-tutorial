package controllers

import (
	"fmt"
	// "github.com/revel/revel"
	"myapp/app/models"
)

type BookModel struct {
	GormController
}

func (b BookModel) insertBook(name, author string) bool {
	new_book := models.Book{Name: name, Author: author}
	b.DB.Create(&new_book)
	return b.DB.NewRecord(new_book)
}

func (b BookModel) getBooks() []*models.Book {
	var count = 0
	// var output = ""
	books := []*models.Book{}
	// books.Id = 1
	b.DB.Find(&books).Count(&count)
	// c.DB.Last(&books).Count(&count)
	// c.DB.First(&books).Count(&count)
	if count == 0 {
		fmt.Println("該当レコードなし")
		// } else {
		// 	for i := 0; i < len(books); i++ {
		// 		output += "\nbook:" + fmt.Sprint(books[i].Id) + " =id/ " + books[i].Name + " / " + books[i].Author
		// 	}
	}
	defer b.DB.Close()
	return books
}
