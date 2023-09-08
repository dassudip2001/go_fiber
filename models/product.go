package models

import "time"

type Product struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Serial    string `json:"serial"`
	CreatedAt time.Time
}
