package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).ConnectDB, revel.BEFORE)
}
