package models

import (
	"fmt"
	"github.com/h-tko/echo-base/libraries"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitDB ...
//
// DB接続の初期化処理.
func InitDB() error {
	conf, err := libraries.GetConfig()

	if err != nil {
		return err
	}

	host := conf.GetString("database.host")
	user := conf.GetString("database.user")
	password := conf.GetString("database.password")
	port := conf.GetString("database.port")
	dbname := conf.GetString("database.dbname")
	dbsource, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmodel=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		return err
	}

	db = dbsource
	db.LogMode(true)

	return nil
}

// CloseDB ...
//
// DB切断処理.
//
// InitDB呼び出し直後にdefer登録すること.
func CloseDB() {
	db.Close()
}
