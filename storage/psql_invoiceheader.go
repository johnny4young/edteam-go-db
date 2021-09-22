package storage

import (
	"database/sql"
	"fmt"

	"github.com/johnny4young/edteam-go-db/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(25) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`
)

// PsqlInvoiceHeader used for working with postgress - invoiceheader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader returns a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implements the interface invoiceheader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
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

// CreateTx implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
