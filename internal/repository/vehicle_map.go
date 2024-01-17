package repository

import "app/internal"

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

//Add is a method that adds a vehicle //Exercise 1 POST /vehicles
func (r *VehicleMap) Add(v internal.Vehicle) (err error){
	
	// check if vehicle already exists
	_, ok := r.db[v.Id]

	if ok {
		return internal.ErrorVehicleAlreadyExists
	}

	// add vehicle
	r.db[v.Id] = v

	return 
}

//Search vehicles by color and year //Exercise 2 GET /vehicles/color/{color}/year/{year}
func (r *VehicleMap) SearchByColorAndYear(color string, year int) (v []internal.Vehicle, err error) {

	for _, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v = append(v, value)
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrorVehiclesNotFound
	}

	return v, nil
}

//Search vehicles by brand and year range //Exercise 3 GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (r *VehicleMap) SearchByBrand(brand string, start_year int, end_year int) (v []internal.Vehicle, err error) {

	for _, value := range r.db{
		if value.Brand == brand && value.FabricationYear >= start_year && value.FabricationYear <= end_year {
			v = append(v, value)
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrorVehiclesNotFound
	}

	return v, nil
}

//Get average speed by brand //Exercise 4 GET /vehicles/average_speed/brand/{brand}
func (r *VehicleMap) GetAverageSpeedByBrand(brand string) (avgSpeed float64, err error) {

	var size float64

	for _, value := range r.db {
		if value.Brand == brand {
			avgSpeed += value.MaxSpeed
			size++
		}
	}

	if size == 0 {
		return 0, internal.ErrorVehiclesNotFound
	}

	return avgSpeed / size, nil
}

//Add multiple vehicles //Exercise 5 POST /vehicles/batch
func (r *VehicleMap) AddMultiple(vehicles []internal.Vehicle) (err error) {
	for _, value := range vehicles {

		// check if vehicle already exists
		_, ok := r.db[value.Id]

		if ok {
			return internal.ErrorVehicleAlreadyExists
		}

		r.db[value.Id] = value

	}
	return nil
}

//Update max speed by id //Exercise 6 PUT /vehicles/{id}/update_speed
func (r *VehicleMap) UpdateMaxSpeedById(id int, maxSpeed float64) (err error) {

	if entry, ok := r.db[id]; ok {

		entry.MaxSpeed = maxSpeed
		r.db[id] = entry
		return nil

	}

	return internal.ErrorVehicleNotFound
}

//Search vehicles by fuel_type //Exercise 7 GET /vehicles/fuel_type/{fuel_type}
func (r *VehicleMap) GetVehiclesByFuelType(fuelType string) (v []internal.Vehicle, err error) {

	for _, value := range r.db {

		if value.FuelType == fuelType {

			v = append(v, value)

		}

	}

	if len(v) == 0 {
		return nil, internal.ErrorVehiclesNotFound
	}

	return v, nil
}

//Delete a vehicle by id //Exercise 8 DELETE /vehicles/{id}
func (r *VehicleMap) DeleteById(id int) (err error) {

	if _, ok := r.db[id]; ok {

		delete(r.db, id)
		return nil

	}

	return internal.ErrorVehicleNotFound
}

//Search vehicles by transmission type //Exercise 9 GET /vehicles/transmission/{transmission}
func (r *VehicleMap) GetVehiclesByTransmission(transmission string) (v []internal.Vehicle, err error) {

	for _, value := range r.db {

		if value.Transmission == transmission {
			v = append(v, value)
		}

	}

	if len(v) == 0 {
		return nil, internal.ErrorVehiclesNotFound
	}

	return v, nil
}

//Update fuel type by id //Exercise 10 PUT /vehicles/{id}/update_fuel
func (r *VehicleMap) UpdateFuelTypeById(id int, fuelType string) (err error) {

	if entry, ok := r.db[id]; ok {

		entry.FuelType = fuelType
		r.db[id] = entry
		return nil

	}

	return internal.ErrorVehicleNotFound
}

//Get average capacity of people by brand //Exercise 11 GET /vehicles/average_capacity/brand/{brand}
func (r *VehicleMap) GetAverageCapacityByBrand(brand string) (avgCapacity int, err error) {

	var size int

	for _, value := range r.db {

		if value.Brand == brand {
			avgCapacity += value.Capacity
			size++
		}

	}

	if size == 0 {

		return 0, internal.ErrorVehiclesNotFound

	}

	return avgCapacity / size, nil
}

//Search vehicles by a range of dimensions of length and width //Exercise 12 GET /vehicles/dimensions?length={min_length}-{max_length}&width={min_width}-{max_width}
func (r *VehicleMap) GetVehiclesByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v []internal.Vehicle, err error) {

	for _, value := range r.db {


		if value.Dimensions.Length >= minLength && value.Dimensions.Length <= maxLength && value.Dimensions.Width >= minWidth && value.Dimensions.Width <= maxWidth {
			v = append(v, value)
		}

	}

	if len(v) == 0 {
		return nil, internal.ErrorVehiclesNotFound
	}

	return v, nil

}

//Search vehicles based by a range of weight //Exercise 13 GET /vehicles/weight?min_weight={min_weight}&max_weight={max_weight}
func (r *VehicleMap) GetVehiclesByWeight(minWeight float64, maxWeight float64) (v []internal.Vehicle, err error) {

	for _, value := range r.db {

		if value.Weight >= minWeight && value.Weight <= maxWeight {
			v = append(v, value)
		}

	}

	if len(v) == 0 {
		return nil, internal.ErrorVehiclesNotFound
	}

	return v, nil
}