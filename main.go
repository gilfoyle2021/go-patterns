package main

import (
	"github.com/SongYintao/go-patterns/pkg/creational/builder"
)

func buildTest() {
	manufactor := builder.Manufactor{}
	carBuilder := &builder.SmallCarBuilder{
		Car: &builder.SmallCar{},
	}
	manufactor.SetBuilder(carBuilder)
	manufactor.Construct()

	simpleCar := carBuilder.Build()
	simpleCar.Drive()
	simpleCar.Stop()
}

func main() {
	buildTest()
}
