package model

func (b *Employee) TableName() string {
	return "employee"
}

type Employee struct {
	EmpNo     int    `json:"emp_no"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Education string `json:"education"`
	DeptNo    int    `json:"dept_no"`
	Position  string `json:"position"`
}

type Department struct {
	DeptNo   int    `json:"dept_no"`
	DeptName string `json:"dept_name"`
}

type Bird struct {
	Id      int    `json:"id"`
	Color   string `json:"color"`
	Typee   string `json:"typee"`
	SteelId int    `json:"steel_id"`
	SinkId  int    `json:"sink_id"`
}
