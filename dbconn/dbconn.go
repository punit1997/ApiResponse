package dbconn

import (
	"github.com/jinzhu/gorm"
	"github.com/punit1997/ApiResponse/models/eecom"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error
func Conn(){
	DB, err = gorm.Open("mysql", "root:12345678@(localhost)/apiresponse?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
}

func Migrate(){
	DB.AutoMigrate(&eecom.JobOrderRequest{})
	DB.AutoMigrate(&eecom.JobOrderResponse{})
}
