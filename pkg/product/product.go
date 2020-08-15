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

// Models slice of Model
type Models []*Model

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

// Service of product
type Service struct {
	storage Storage
}

// NewService returns a new pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is being used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
