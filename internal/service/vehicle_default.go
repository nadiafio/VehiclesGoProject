package service

import (
	"app/internal"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Add is a method that adds a vehicle //Exercise 1 POST /vehicles
func (s *VehicleDefault) Add(v internal.Vehicle) (err error) {
	err = s.rp.Add(v)

	if err != nil{
		switch err {

			case internal.ErrorVehicleAlreadyExists:
				err = fmt.Errorf("%w: id", internal.ErrorVehicleAlreadyExists)

			}

			return
		}
	
	return
}

// Search vehicles by color and year //Exercise 2 GET /vehicles/color/{color}/year/{year}
func (s *VehicleDefault) SearchByColorAndYear(color string, year int) (v []internal.Vehicle, err error) {
	v, err = s.rp.SearchByColorAndYear(color, year)

	if err != nil {

		err = fmt.Errorf("%w", internal.ErrorVehiclesNotFound)
		return

	}

	return
}

// Search vehicles by brand and year range //Exercise 3 GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (s *VehicleDefault) SearchByBrand(brand string, start_year int, end_year int) (v []internal.Vehicle, err error) {

	v, err = s.rp.SearchByBrand(brand, start_year, end_year)

	if err != nil {

		err = fmt.Errorf("%w", internal.ErrorVehiclesNotFound)
		return

	}

	return
}

// Get average speed by brand //Exercise 4 GET /vehicles/average_speed/brand/{brand}
func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (avgSpeed float64, err error) {

	avgSpeed, err = s.rp.GetAverageSpeedByBrand(brand)

	if err != nil {
		err = fmt.Errorf("%w", internal.ErrorVehiclesNotFound)
		return
	}
	
	return
}