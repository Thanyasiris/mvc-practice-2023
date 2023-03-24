package controller

import (
	"math/rand"
	"mvc_structure/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get emp
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

// create employee
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

// edit employee
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

// delete employee
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

// create bird
func CreateBird(c *gin.Context) {
	var err error
	var RandomColor [3]int //[0]blue [1]green [2]yellow
	var color = [3]string{"blue", "green", "yellow"}
	var typ = [2]string{"RET", "DM"}
	var num = 0
	for i := 0; i < 3; i++ {
		RandomColor[i] = rand.Intn(81-40) + 40
		num += RandomColor[i]
	}
	member := make([]model.Bird, int(num))

	k := 0
	cage := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < RandomColor[i]; j++ {
			member[k].Id = k + 1
			member[k].Color = color[i]
			ty := rand.Intn(2-0) + 0
			member[k].Typee = typ[ty]
			member[k].SteelId = cage
			if (j+1)%7 == 0 {
				cage++
			}
			member[k].SinkId = 0
			err = model.CreateBird(&member[k])
			k++
		}
		cage++
	}

	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, member)
		return
	}

}
