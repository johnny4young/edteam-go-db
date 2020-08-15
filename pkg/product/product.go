package product

import (
	"time"
)

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreateAt     time.Time
	UpdatedAt    time.Time
}
