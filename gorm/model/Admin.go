package model

type Admin struct {
	Id           int
	Name         string
	Password     string
	Sex          string
	Phonenum     int
	Jurisdiction string
}

func (*Admin) TableName() string {
	return "admin"
}
