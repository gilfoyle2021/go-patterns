package builder

import "fmt"

//Color car clolor
type Color string

const (
	//Red red color
	Red Color = "red"
	//Blue blue color
	Blue = "blue"
)

//Speed car speed
type Speed int

const (
	//MPH miles per hour
	MPH Speed = 1
	//KPH kellometers per hour
	KPH = 2
)

//Wheel wheel type
type Wheel string

const (
	//SteelWheel steel Wheel
	SteelWheel Wheel = "steelWheel"
	//SportsWheel Sports Wheel
	SportsWheel = "sportWheel"
)

//Builder construct builder
type Builder interface {
	Color(Color) Builder
	Wheels(Wheel) Builder
	TopSpeed(Speed) Builder
	Build() Car
}

//Car car interface
type Car interface {
	Drive() error
	Stop() error
}

//SmallCar smalll car
type SmallCar struct {
	Color Color
	Speed Speed
	Wheel Wheel
}

// Drive drive method
func (sc *SmallCar) Drive() error {
	fmt.Println("Drive a small car...")
	return nil
}

// Stop stop method
func (sc *SmallCar) Stop() error {
	fmt.Println("Stop a small car...")
	return nil
}

// SmallCarBuilder small car builder
type SmallCarBuilder struct {
	Car *SmallCar
}

// Color car's color
func (scb *SmallCarBuilder) Color(color Color) Builder {
	scb.Car.Color = color
	return scb
}

// Wheels car's wheel
func (scb *SmallCarBuilder) Wheels(wheel Wheel) Builder {
	scb.Car.Wheel = wheel
	return scb
}

// TopSpeed top speed
func (scb *SmallCarBuilder) TopSpeed(speed Speed) Builder {
	scb.Car.Speed = speed
	return scb
}

// Build car build
func (scb *SmallCarBuilder) Build() Car {
	return scb.Car
}

//Manufactor Car manufacter
type Manufactor struct {
	builder Builder
}

//SetBuilder set real Car builder
func (carManu *Manufactor) SetBuilder(builder Builder) {
	carManu.builder = builder
}

//Construct build Builder by parameters
func (carManu *Manufactor) Construct() {
	//set car build according to info
	carManu.builder.Color(Red).TopSpeed(MPH).Wheels(SteelWheel)
}
