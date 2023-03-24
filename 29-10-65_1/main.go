package main

import (
	"fmt"

	"mvc_structure/config"
	Routes "mvc_structure/controller"
	"mvc_structure/model"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	//swagger
	fmt.Println("Run Swagger : http://localhost:8080/swagger/index.html")
	fmt.Println()
	//database configure
	config.DB, err = gorm.Open("postgres", config.DBStr(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&model.Employee{})
	r := Routes.SetUpRouter()
	//running
	r.Run()
}
