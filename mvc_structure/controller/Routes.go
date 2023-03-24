package controller

import (
	_ "mvc_structure/view"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter() *gin.Engine {
	//Set Up Router ... Configure routes
	r := gin.Default()
	r.GET("/getAllEmployee", GetAllEmployee)
	r.POST("/createEmployee", CreateEmployee)
	r.POST("/editEmployee", EditEmployee)
	r.POST("/deleteEmployee", DeleteEmployee)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
