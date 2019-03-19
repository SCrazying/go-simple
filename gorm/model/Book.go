package model

import "time"

type Book struct {
	Id             int
	Name           string
	Sn             string
	Author         string
	Press          string
	PublishingTime time.Time
	Price          float32
}

func (*Book) TableName() string {
	return "book"
}
