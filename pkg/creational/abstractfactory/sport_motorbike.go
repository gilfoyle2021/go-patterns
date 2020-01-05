package abstractfactory

type SportMotorbike struct {
}

func (sm *SportMotorbike) NumSeats() int {
	return 2
}
func (sm *SportMotorbike) NumWheels() int {
	return 4
}

func (sm *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
