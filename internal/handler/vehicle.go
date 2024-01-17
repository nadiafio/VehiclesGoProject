package handler

import (
	"app/internal"
	"strconv"
	// "app/platform/web"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type Message struct{
	Message string
	Data any
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Add() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			
			return
		}
		
		var bodyMap map[string]any
		
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			
			return
		}
		
		if err := ValidateKeyExistance(bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, err.Error())
			
			return
		}
		
		var body VehicleJSON

		if err := json.Unmarshal(bytes, &body); err != nil{
			response.Text(w, http.StatusBadRequest, "invalid request body")
			
			return
		}

		vehicle := internal.Vehicle{
			Id : body.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		if err := h.sv.Add(vehicle); err != nil {

			response.Text(w, http.StatusConflict, err.Error())

			return

		}

		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Dimensions.Height,
			Length:          vehicle.Dimensions.Length,
			Width:           vehicle.Dimensions.Width,
		}

		response.JSON(w, http.StatusOK, &Message{
			Message: "movie created successfully",
			Data:    data,
		})

	}
}

func ValidateKeyExistance(body map[string]any) error {

	keys := []string{"id", "brand", "model", "registration", "color", "year", "passengers", "max_speed", "fuel_type", "transmission", "weight", "height", "width"}

	for _, key := range keys {

		if _, ok := body[key]; !ok {

			return fmt.Errorf("key %s not found", key)
			
		}

	}
	return nil
}

func (h *VehicleDefault) SearchByColorAndYear() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r,"year"))

		if err != nil{
			response.Text(w, http.StatusBadRequest, "invalid year")
			return
		}

		v, err := h.sv.SearchByColorAndYear(color, year)

		if err != nil {
			response.Text(w, http.StatusConflict, err.Error())
			return
		}

		vehicles := []VehicleJSON{}

		for _, value := range v{

			vehicles = append(vehicles, VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Dimensions.Height,
				Length:          value.Dimensions.Length,
				Width:           value.Dimensions.Width,
			})
			
		}

		response.JSON(w, http.StatusOK, &Message{
			Message: "movies found successfully",
			Data:    vehicles,
		})

	}
}

func (h *VehicleDefault) SearchByBrand() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		brand := chi.URLParam(r, "brand")
		start, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		end, err := strconv.Atoi(chi.URLParam(r, "end_year"))

		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid year")
			return
		}

		v, err := h.sv.SearchByBrand(brand, start, end)

		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		vehicles := []VehicleJSON{}

		for _, value := range v{
			vehicles = append(vehicles, VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Dimensions.Height,
				Length:          value.Dimensions.Length,
				Width:           value.Dimensions.Width,
			})
		}

		response.JSON(w, http.StatusOK, &Message{
			Message: "vehicles found successfully",
			Data:    vehicles,
		})


	}
}

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		brand := chi.URLParam(r, "brand")

		speed, err := h.sv.GetAverageSpeedByBrand(brand)

		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}
		
		response.JSON(w, http.StatusOK, &Message{
			Message: "average speed found successfully",
			Data:    speed,
		})

	}
}