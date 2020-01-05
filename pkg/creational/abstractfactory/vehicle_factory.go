package abstractfactory

import "errors"

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarVehicle       = 1
	Motorbikevehicle = 2
)

func BuildVehicleFactory(vehicleType int) (VehicleFactory, error) {

	switch vehicleType {
	case CarVehicle:
		return new(CarFactory), nil
	case Motorbikevehicle:
		return new(MotorBikefactory), nil
	default:
		return nil, errors.New("Not support this type")
	}
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct {
}

func (cf *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("Not support this type")
	}
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorBikefactory struct {
}

func (cf *MotorBikefactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New("Not support this type")
	}
}
