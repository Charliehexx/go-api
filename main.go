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
	app.PUT("/car/:id/update", func(ctx *gofr.Context) (interface{}, error) {
		// Extract car ID from the route parameters
		carID := ctx.PathParam("id")

		var carDetails Car
		var updateRequest struct {
			RepairStatus string `json:"repair_status"`
		}

		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&carDetails); err != nil {

			return nil, err
		}

		// Update the car entry in the database with the new repair status
		_, err := ctx.DB().ExecContext(ctx, "UPDATE cars SET repair_status = ? WHERE id = ?", updateRequest.RepairStatus, carID)

		if err != nil {
			// Handle the error, e.g., return an HTTP 500 for a server error
			return nil, err
		}

		return "Car repair status updated successfully", nil
	})

	app.Start()
}
