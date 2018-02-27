package abstract_factory

type SportMotorbike struct{}

func (*SportMotorbike) NumDoors() int {
	return 5
}

func (*SportMotorbike) NumWheels() int {
	return 4
}

func (*SportMotorbike) NumSeats() int {
	return 5
}

func (*SportMotorbike) GetMotorBikeType() int {
	return SportMotorbikeType
}
