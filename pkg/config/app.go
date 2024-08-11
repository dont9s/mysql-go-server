package config

import(

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (

	db *gorm.DB
)

func Connect(){

	createdDb, err := gorm.Open("mysql", "dont9s:dont9s/simplerest?charset=utf8&parseTime=True&loc=local")

	if err != nil{
		panic(err)

	}
	db = createdDb
}


func GetDB() *gorm.DB{

	return db
}



