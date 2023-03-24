package model

import (
	"mvc_structure/config"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllEmp
func ListAllEmp(emp *[]Employee) (err error) {
	if err = config.DB.Find(&emp).Error; err != nil {
		return
	}

	return
}

// CreateEmp
func CreateEmp(req *Employee) (err error) {
	if err = config.DB.Table("employee").
		Create(&req).Error; err != nil {
		return
	}

	return
}

// EditEmp
func EditEmp(req *Employee) (err error) {
	if err = config.DB.Table("employee").
		Where("emp_no = ?", req.EmpNo).
		Update(map[string]interface{}{
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"education":  req.Education,
			"position":   req.Position,
			"dept_no":    req.DeptNo,
		}).
		Error; err != nil {
		return
	}
	return
}

// DeleteEmp
func DeleteEmp(req *Employee) (err error) {
	if err = config.DB.Table("employee").
		Where("emp_no = ?", req.EmpNo).
		Delete(&req).Error; err != nil {
		return
	}
	return
}
