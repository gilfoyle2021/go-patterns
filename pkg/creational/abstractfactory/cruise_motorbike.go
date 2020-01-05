package abstractfactory

type CruiseMotorbike struct {
}

func (sm *CruiseMotorbike) NumSeats() int {
	return 2
}
func (sm *CruiseMotorbike) NumWheels() int {
	return 4
}

func (sm *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
