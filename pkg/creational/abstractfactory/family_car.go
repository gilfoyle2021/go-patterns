package abstractfactory

type FamilyCar struct {
}

func (fc *FamilyCar) NumSeats() int {
	return 41
}

func (fc *FamilyCar) NumDoors() int {
	return 41
}

func (fc *FamilyCar) NumWheels() int {
	return 61
}
