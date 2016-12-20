package dealership

import "fmt"

func Purchase(make string, model string) *Car {
	c := &Car{
		make:  make,
		model: model,
	}

	fmt.Println("Congratulations! You just got yourself a " + c.Make() + " " + c.Model())
	return c
}

func Sell(c *Car) {
	fmt.Println("Your " + c.Make() + " " + c.Model() + " is sold!")
}

type Car struct {
	make  string
	model string
}

func (c *Car) Make() string {
	return c.make
}

func (c *Car) Model() string {
	return c.model
}
