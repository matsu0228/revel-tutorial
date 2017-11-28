package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
	"strings"
)

func getConfigParam(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
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

var DB *gorm.DB

func GormConnect() {
	connectionString := getConnectionString()
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	// return db
	db.DB()
	DB = db
}
