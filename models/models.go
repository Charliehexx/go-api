package models

import (
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
