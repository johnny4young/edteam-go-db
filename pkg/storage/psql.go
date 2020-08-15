package storage

import (
	"database/sql"
)

// PsqlProduct used for working with postgress - produc
type PsqlProduct struct {
	db *sql.DB
}
