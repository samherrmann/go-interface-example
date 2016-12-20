package joemotors

import "fmt"

type Vehicle interface {
	Make() string
	Model() string
}

func Service(v Vehicle) {
	fmt.Println("Joe Motos is servicing your " + v.Make() + " " + v.Model())
}
