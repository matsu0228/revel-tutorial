package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app/models"
)

type App struct {
	GormController
}

func (c App) getBooks() string {
	var count = 0
	var output = ""

	// insert
	// ---------------------------
	new_book := models.Book{Name: "KimiNoNaha?", Author: "Shinkai"}
	// if c.DB.NewRecord(new_book) // => returns `true` as primary key is blank
	c.DB.Create(&new_book)

	// select
	// ---------------------------
	books := []*models.Book{}
	// books.Id = 1
	c.DB.Find(&books).Count(&count)
	// c.DB.Last(&books).Count(&count)
	// c.DB.First(&books).Count(&count)
	if count == 0 {
		fmt.Println("該当レコードなし")
	} else {
		for i := 0; i < len(books); i++ {
			output += "\nbook:" + fmt.Sprint(books[i].Id) + " =id/ " + books[i].Name + " / " + books[i].Author
		}
	}
	defer c.DB.Close()
	return output
}

// ----------------------------------------------------------------------------
func (c App) Index() revel.Result {
	// return c.Render()
	// connectionString := getConnectionString()
	book_data := c.getBooks()
	greeting := "Aloha World/" + book_data
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}
