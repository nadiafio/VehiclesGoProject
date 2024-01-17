package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add is a method that adds a vehicle
	Add(v Vehicle) (err error)
	// Search vehicles by color and year
	SearchByColorAndYear(color string, year int) (v []Vehicle, err error)
	// Search vehicles by brand and year range
	SearchByBrand(brand string, start_year int, end_year int) (v []Vehicle, err error)
	// Get average speed by brand
	GetAverageSpeedByBrand(brand string) (avgSpeed float64, err error)
	// Add multiple vehicles
	// AddMultiple(vehicles []Vehicle) (err error)
	// // Update max speed by id
	// UpdateMaxSpeedById(id int, maxSpeed float64) (err error)
	// // Search vehicles by fuel_type
	// GetVehiclesByFuelType(fuelType string) (v []Vehicle, err error)
	// // Delete a vehicle by id
	// DeleteById(id int) (err error)
	// // Search vehicles by transmission type
	// GetVehiclesByTransmission(transmission string) (v []Vehicle, err error)
	// // Update fuel type by id
	// UpdateFuelTypeById(id int, fuelType string) (err error)
	// // Get average capacity of people by brand
	// GetAverageCapacityByBrand(brand string) (avgCapacity int, err error)
}