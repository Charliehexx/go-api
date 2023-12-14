package main

import (
	"encoding/json"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"time"
)

type Car struct {
	Id           int       `json:"id"`
	LicensePlate string    `json:"license_plate"`
	Model        string    `json:"model"`
	Color        string    `json:"color"`
	RepairStatus string    `json:"repair_status"`
	EntryTime    time.Time `json:"entry_time"`
}

func main() {
	app := gofr.New()
	//post route --- for posting a car in the garage (car added in database)
	app.POST("/car/enter", Createcar)
	//get route--- get all the cars till now saved on the database
	app.GET("/car", Getcars)
	//update route--- can change the repair status of the car by passing specific endpoint
	app.PUT("/car/update/{id}", Updatecar)
	//delete route --delete a particular car by hitting on the specific id endpoint
	app.DELETE("/car/delete/{id}", Deletecar)
	app.Start()
}

func Createcar(ctx *gofr.Context) (interface{}, error) {
	var carDetails Car
	//decoding the json data from the body
	decoder := json.NewDecoder(ctx.Request().Body)
	if err := decoder.Decode(&carDetails); err != nil {
		return nil, errors.MissingParam{Param: []string{"license_plate,model,color,repair_status"}} // error check for fields cannot be empty while making post request
	}
	//license plate cannot be empty
	if carDetails.LicensePlate == "" {
		return nil, errors.MissingParam{Param: []string{"license_plate"}}
	}

	// Check if the number plate is unique
	if isNumberPlateExists(ctx, carDetails.LicensePlate) {
		return nil, errors.InvalidParam{Param: []string{"license_plate should be different"}}
	}
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO cars (license_plate, model, color,repair_status) VALUES (?, ?, ?,?)",
		carDetails.LicensePlate, carDetails.Model, carDetails.Color, carDetails.RepairStatus)
	if err != nil {
		return nil, err
	}

	return "Car entered the garage successfully", errors.EntityAlreadyExists{}
}

func Getcars(ctx *gofr.Context) (interface{}, error) {
	var cars []Car
	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM cars")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.Id, &car.LicensePlate, &car.Color, &car.Model, &car.RepairStatus, &car.EntryTime); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
func Updatecar(ctx *gofr.Context) (interface{}, error) {
	carID := ctx.PathParam("id")
	var updateRequest struct {
		RepairStatus string `json:"repair_status"`
	}
	decoder := json.NewDecoder(ctx.Request().Body)
	if err := decoder.Decode(&updateRequest); err != nil {
		return nil, err
	}
	_, err := ctx.DB().ExecContext(ctx, "UPDATE cars SET repair_status = ? WHERE id = ?", updateRequest.RepairStatus, carID)
	if err != nil {
		return nil, err
	}
	return "Car repair status updated successfully", nil
}

func Deletecar(ctx *gofr.Context) (interface{}, error) {
	carID := ctx.PathParam("id")
	if !isCarIDExists(ctx, carID) {
		return nil, errors.EntityNotFound{Entity: "car with this ID"}
	}
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM cars WHERE id = ?", carID)
	if err != nil {
		return nil, err
	}
	return "Car deleted successfully", nil
}

// ---------------------HELPER FUNCTIONS-------------------------
// helper function for post route
func isNumberPlateExists(ctx *gofr.Context, licensePlate string) bool {
	var count int
	// Execute a query to count cars with the given number plate
	err := ctx.DB().QueryRowContext(ctx, "SELECT COUNT(*) FROM cars WHERE license_plate = ?", licensePlate).Scan(&count)
	if err != nil {
		// Handle the error (e.g., log or return false)
		return false
	}
	// If count is greater than 0, the number plate already exists
	return count > 0
}

// helper function for delete route to check if carID already present in database
func isCarIDExists(ctx *gofr.Context, carID string) bool {
	var count int
	// Execute a query to count cars with the given ID
	err := ctx.DB().QueryRowContext(ctx, "SELECT COUNT(*) FROM cars WHERE id = ?", carID).Scan(&count)
	if err != nil {
		// Handle the error (e.g., log or return false)
		return false
	}
	// If count is greater than 0, the record with the specified ID exists
	return count > 0
}
