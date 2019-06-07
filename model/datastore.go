package model

import (
	"fmt"

	config "github.com/PandeKaustubhS/microservice-todo/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection(conf config.DBConfig) (*gorm.DB, error) {
	// MYSQL Connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	return gorm.Open("mysql", connectionString)
}
