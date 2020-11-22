package models

type Place struct {
	Country string `orm:"column(country);size(200);null"`
	City    string `orm:"column(city);size(200);null"`
	Telcode int    `orm:"column(telcode);null"`
}
