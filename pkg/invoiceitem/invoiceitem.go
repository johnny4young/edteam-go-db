package invoiceitem

import (
	"time"
)

// Model of invoiceitem
type Model struct {
	ID               uint
	InvoiceHeadderID uint
	ProductID        uint
	CreatedAt        time.Time
	UupdatedAt       time.Time
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

// Service of invoice item
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
