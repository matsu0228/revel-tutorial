package controllers

import (
	"fmt"
	"github.com/revel/revel"
	// "myapp/app/models"
)

type App struct {
	// GormController
	BookModel
}

func (c App) Index() revel.Result {
	result := c.insertBook("name is this", "author is me")
	if !result {
		fmt.Println("insert err")
	}
	book_data := c.getBooks()
	return c.RenderJSON(book_data)
	// greeting := "Aloha World/" + book_data
	// return c.Render(greeting)
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
