package main

import (
	"encoding/json"
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
	app.POST("/car/enter", func(ctx *gofr.Context) (interface{}, error) {
		var carDetails Car
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&carDetails); err != nil {
			return nil, err
		}
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO cars (license_plate, model, color,repair_status) VALUES (?, ?, ?,?)",
			carDetails.LicensePlate, carDetails.Model, carDetails.Color, carDetails.RepairStatus)
		if err != nil {
			return nil, err
		}
		return "Car entered the garage successfully", nil
	})
	//get route--- get all the cars till now saved on the database
	app.GET("/car", func(ctx *gofr.Context) (interface{}, error) {
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
	})
	//update route--- can change the repair status of the car by passing specific endpoint
	app.PUT("/car/update/{id}", func(ctx *gofr.Context) (interface{}, error) {
		carID := ctx.PathParam("id")
		//fmt.Println(carID)
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
	})
	//delete route --delete a particular car by hitting on the specific id endpoint
	app.DELETE("/car/delete/{id}", func(ctx *gofr.Context) (interface{}, error) {
		carID := ctx.PathParam("id")
		_, err := ctx.DB().ExecContext(ctx, "DELETE FROM cars WHERE id = ?", carID)
		if err != nil {
			return nil, err
		}
		return "Car with ID %s deleted successfully", nil
	})
	app.Start()
}
