package database

import (
	"database/sql"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

//go:generate counterfeiter -o fake_db.go -pkg ${GOPACKAGE} -fake-name FakeDB . DB

// DB describes the database interface used by RequestScope.
type DB interface {
	dbx.Builder
	Begin() (Tx, error)
	Close() error
	DriverName() string
	Original() *dbx.DB
	Instance() *sql.DB
	SetLogFunc(dbx.LogFunc)
}

type dbxWrapper struct {
	dbx.DB
}

func (d *dbxWrapper) Begin() (Tx, error) {
	tx, err := d.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &txWrapper{
		Tx: tx,
		db: &d.DB,
	}, nil
}

func (d *dbxWrapper) Original() *dbx.DB {
	return &d.DB
}

func (d *dbxWrapper) Select(cols ...string) *dbx.SelectQuery {
	sq := d.DB.Select(cols...)
	sq.FieldMapper = d.DB.FieldMapper
	return sq
}

func (d *dbxWrapper) Instance() *sql.DB {
	return d.DB.DB()
}

func (d *dbxWrapper) SetLogFunc(fn dbx.LogFunc) {
	d.DB.LogFunc = fn
}

// New creates a new DB out of a dbx.DB.
func New(dbxDB *dbx.DB) DB {
	return &dbxWrapper{
		*dbxDB,
	}
}
