package mysql

import (
	"database/sql"
	"time"

	"bitbucket.org/xeoncross/godiapp"
)

// Several design choices here
//
// Wrap sql.DB and sql.Tx
// https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
//
// HTTP handler usage:
// https://gist.github.com/tsenart/5fc18c659814c078378d
//
// DI:
// https://www.alexedwards.net/blog/organising-database-access#using-an-interface

// type Datastore interface {
// 	AllSends() ([]*Send, error)
// }

type DB struct {
	*sql.DB
}

// type Tx struct {
// 	*sql.Tx
// }

func NewDB(dataSourceName string) (*DB, error) {
	c, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = c.Ping(); err != nil {
		return nil, err
	}

	// 1. If any other process tries to connect to the database, Go could starve
	// them out if we don't limit the max connections it can hold in it's pool.
	// 2. Lots of goroutines (http handlers)? Consider the following MySQL error:
	// "Error 1461: Can't create more than max_prepared_stmt_count statements"
	// 3. [My|Postgre|MS]SQL can only perform so many actions at once anyway, so
	// having more connections open doesn't really buy us anything.
	c.SetMaxOpenConns(100)

	// Release connections we aren't using anymore
	c.SetConnMaxLifetime(time.Second * 120)

	return &DB{c}, nil
}

func (db *DB) GetUsers() ([]*godiapp.User, error) {
	rows, err := db.Query("SELECT id, email FROM user LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ss := make([]*godiapp.User, 0)
	for rows.Next() {
		s := new(godiapp.User)
		err := rows.Scan(&s.ID, &s.Email)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ss, nil
}

// // Begin starts an returns a new transaction.
// func (db *db) Begin() (*Tx, error) {
// 	tx, err := db.DB.Begin()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Tx{tx}, nil
// }

// // NullTime represents a time.Time that may be null. NullTime implements the
// // sql.Scanner interface so it can be used as a scan destination, similar to
// // sql.NullString. https://github.com/lib/pq/blob/master/encode.go#L583
// type NullTime struct {
// 	Time  time.Time
// 	Valid bool // Valid is true if Time is not NULL
// }
//
// // Scan implements the Scanner interface.
// func (nt *NullTime) Scan(value interface{}) error {
// 	nt.Time, nt.Valid = value.(time.Time)
// 	return nil
// }
//
// // Value implements the driver Valuer interface.
// func (nt NullTime) Value() (driver.Value, error) {
// 	if !nt.Valid {
// 		return nil, nil
// 	}
// 	return nt.Time, nil
// }
