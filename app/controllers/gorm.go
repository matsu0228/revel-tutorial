package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	r "github.com/revel/revel"
	"strings"
)

type GormController struct {
	*r.Controller
	DB *gorm.DB
}

func getConfigParam(param string, defaultValue string) string {
	p, found := r.Config.String(param)
	if !found {
		if defaultValue == "" {
			r.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getConfigParam("db.host", "")
	port := getConfigParam("db.port", "3306")
	user := getConfigParam("db.user", "")
	pass := getConfigParam("db.password", "")
	dbname := getConfigParam("db.name", "")
	protocol := getConfigParam("db.protocol", "tcp")
	dbargs := getConfigParam("db.args", " ")
	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

var Gdb *gorm.DB

func InitDB() {
	r.Config.LoadContext("db.conf", ConfPaths)
	// r.LoadConfig("db.conf")
	// r.ConfPaths = []string{"/myapp/conf/app.conf", "/myapp/conf/db.conf"}
	var err error
	connectionString := getConnectionString()
	Gdb, err = gorm.Open("mysql", connectionString)
	Gdb.LogMode(true)
	if err != nil {
		panic(err.Error())
	}
}
func (c *GormController) ConnectDB() r.Result {
	c.DB = Gdb
	return nil
}
