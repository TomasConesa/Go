package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
)

// Patron de diseño Factory
func New(vehicle string, time int) (Vehicle, error) {
	switch vehicle {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	}
	return nil, fmt.Errorf("Vehicle '%s' doesn't exist", vehicle)
}

type Car struct {
	Time int
}

// Suponiendo que son 100 km/h
func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

type Motorcycle struct {
	Time int
}

func (m *Motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60)
}

type Truck struct {
	Time int
}

func (t *Truck) Distance() float64 {
	return 70 * (float64(t.Time) / 60)
}
