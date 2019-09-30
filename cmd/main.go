package main

import (
	"github.com/jinzhu/gorm"
	"github.com/punit1997/ApiResponse/dbconn"
	"github.com/punit1997/ApiResponse/routes"
)

var DB *gorm.DB

func main() {


	dbconn.Conn()
	defer dbconn.DB.Close()
	dbconn.Migrate()

	r := routes.InitRoute()
	r.Run(":8081")
}