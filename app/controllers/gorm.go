package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/config"
	r "github.com/revel/revel"
	"path/filepath"
	"strings"
)

type GormController struct {
	*r.Controller
	DB *gorm.DB
}

func getConfigParam(configFile string, section string, param string, defaultValue string) string {
	c, err := config.ReadDefault(configFile)
	if err != nil {
		if defaultValue == "" {
			r.ERROR.Fatal("Cound not find conf file: " + configFile)
		} else {
			return defaultValue
		}
	}
	p, err_read := c.String(section, param)
	if err_read != nil {
		if defaultValue == "" {
			r.ERROR.Fatal("Cound not find section/parameter: " + section + " / " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	// TOOD: set path
	apath, _ := filepath.Abs("./")
	fmt.Println(apath)
	db_conf := "./conf/db.conf"
	host := getConfigParam(db_conf, "database", "db.host", "")
	port := getConfigParam(db_conf, "database", "db.port", "3306")
	user := getConfigParam(db_conf, "database", "db.user", "")
	pass := getConfigParam(db_conf, "database", "db.password", "")
	dbname := getConfigParam(db_conf, "database", "db.name", "")
	protocol := getConfigParam(db_conf, "database", "db.protocol", "tcp")
	dbargs := getConfigParam(db_conf, "database", "db.args", " ")
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
