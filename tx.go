package database

import dbx "github.com/go-ozzo/ozzo-dbx"

//go:generate counterfeiter -o fake_tx.go -pkg ${GOPACKAGE} -fake-name FakeTX . Tx

// Tx is a database transaction.
type Tx interface {
	dbx.Builder
	Commit() error
	Rollback() error
}

type txWrapper struct {
	*dbx.Tx
	db *dbx.DB
}

func (t *txWrapper) Select(cols ...string) *dbx.SelectQuery {
	sq := t.Tx.Select(cols...)
	sq.FieldMapper = t.db.FieldMapper
	return sq
}
