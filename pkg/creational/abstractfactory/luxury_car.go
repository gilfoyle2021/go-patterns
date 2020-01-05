package abstractfactory

type LuxuryCar struct {
}

func (lc *LuxuryCar) NumSeats() int {
	return 4
}

func (lc *LuxuryCar) NumDoors() int {
	return 4
}

func (lc *LuxuryCar) NumWheels() int {
	return 6
}