package models

// import (
// 	"fmt"
// 	// "github.com/revel/revel"
// 	// "myapp/app/controllers"
// 	// 	"regexp"
// 	// 	"time"
// )

type Book struct {
	Id     int    `json:id`
	Name   string `json:name`
	Author string `json:author`
}

// TODO: Validate

// func (book Book) Validate(v *revel.Validation) {
// 	v.Required(booking.User)
//
// 	v.Match(booking.CardNumber, regexp.MustCompile(`\d{16}`)).
// 		Message("Credit card number must be numeric and 16 digits")
//
// 	v.Check(booking.NameOnCard,
// 		revel.Required{},
// 		revel.MinSize{3},
// 		revel.MaxSize{70},
// 	)
// }
