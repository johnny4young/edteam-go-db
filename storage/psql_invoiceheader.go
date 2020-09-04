package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(25) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
)

// PsqlInvoiceheader used for working with postgress - invoiceheader
type PsqlInvoiceheader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader returns a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceheader {
	return &PsqlInvoiceheader{db}
}

// Migrate implements the interface invoiceheader.Storage
func (p *PsqlInvoiceheader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migrations invoice headers done")
	return nil
}
