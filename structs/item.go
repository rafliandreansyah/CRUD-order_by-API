package structs

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint
}
