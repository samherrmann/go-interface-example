package main

import (
	"github.com/samherrmann/go-interface-example/dealership"
	"github.com/samherrmann/go-interface-example/jimautos"
	"github.com/samherrmann/go-interface-example/joemotors"
)

func main() {

	// myCar is of type *Car, which has no affiliation
	// with the Vehicle interface from Joe Motors or the
	// Automobile interface in Jim Autos...
	myCar := dealership.Purchase("BMW", "i8")

	// ...yet simply because they all contain the same
	// methods of the same signature, myCar is portable
	// between all three packages.
	joemotors.Service(myCar)
	jimautos.Fix(myCar)
	dealership.Sell(myCar)

	// Output of this example app:
	//
	// Congratulations! You just got yourself a BMW i8
	// Joe Motos is servicing your BMW i8
	// Jim Autos is fixing your BMW i8
	// Your BMW i8 is sold!
}
