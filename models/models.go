package models

import (
	"fmt"
	"gin-demo/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func init() {
	sec, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatalln(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sec.Key("USER").String(),
		sec.Key("PASSWORD").String(),
		sec.Key("HOST").String(),
		sec.Key("DB_NAME").String(),
	)

	dbType := sec.Key("TYPE").String()
	tablePrefix := sec.Key("TABLE_PREFIX").String()

	//注意这里一定要用外部声明的变量不然访问不到连接
	db, err = gorm.Open(dbType, dsn)
	if err != nil {
		log.Fatalf("%v", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName
	}

}

//func CloseDB() {
//	defer db.Close()
//}
