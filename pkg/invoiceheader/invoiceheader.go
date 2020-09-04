package invoiceheader

import (
	"time"
)

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

// Service of invoice header
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
