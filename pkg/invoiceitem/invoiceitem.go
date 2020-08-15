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
