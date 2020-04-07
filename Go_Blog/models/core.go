package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	//新建存放数据库的文件夹
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed to create datafiles" + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
}
