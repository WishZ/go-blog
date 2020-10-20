package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"re_station/pkg/setting"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                                                error
		dbType, dbName, user, password, host, port, tablePrefix, dbCharset string
		DbMaxOpenConn, DbMaxIdleConn                                       int
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("DB_DRIVER").String()
	dbName = sec.Key("DB_DATABASE").String()
	user = sec.Key("DB_USERNAME").String()
	password = sec.Key("DB_PASSWORD").String()
	host = sec.Key("DB_HOST").String()
	port = sec.Key("DB_PORT").String()
	tablePrefix = sec.Key("DB_PREFIX").String()
	dbCharset = sec.Key("DB_CHARSET").String()
	DbMaxOpenConn = sec.Key("DB_MAX_OPEN_CONN").MustInt(100)
	DbMaxIdleConn = sec.Key("DB_MAX_IDLE_CONN").MustInt(100)
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
		dbCharset))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(DbMaxIdleConn)
	db.DB().SetMaxOpenConns(DbMaxOpenConn)
}
func CloseDB() {
	defer db.Close()
}
