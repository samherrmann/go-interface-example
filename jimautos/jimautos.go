package jimautos

import "fmt"

type Automobile interface {
	Make() string
	Model() string
}

func Fix(a Automobile) {
	fmt.Println("Jim Autos is fixing your " + a.Make() + " " + a.Model())
}
