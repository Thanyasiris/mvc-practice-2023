package controller

import (
	"mvc_structure/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//get emp
func GetAllEmployee(c *gin.Context) {
	var emp []model.Employee
	err := model.ListAllEmp(&emp)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, emp)
		return
	}
}

//create employee
func CreateEmployee(c *gin.Context) {
	var emp model.Employee
	c.BindJSON(&emp)

	err := model.CreateEmp(&emp)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, emp)
		return
	}
}

//edit employee
func EditEmployee(c *gin.Context) {
	var emp model.Employee
	c.BindJSON(&emp)

	err := model.EditEmp(&emp)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, emp)
		return
	}
}

//delete employee
func DeleteEmployee(c *gin.Context) {
	var emp model.Employee
	c.BindJSON(&emp)

	err := model.DeleteEmp(&emp)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, "ลบข้อมูลสำเร็จ")
		return
	}
}
