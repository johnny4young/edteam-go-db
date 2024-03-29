package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	// only need to register de provider
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB   // estructura db gestiona un pool de conexiones activas e inactivas
	once sync.Once // estructura Once que permite ejecutar una única vez (Singleton)
)

// NewPostgresDB create a singleton connection
func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:example@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("connected to postgress")
	})
}

func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:example@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("connected to mySQL")
	})
}

// Pool return an instance unique of db
func Pool() *sql.DB {
	return db
}

// utility to handle null values for string data type
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// utility to handle null values for Time data tipe
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
