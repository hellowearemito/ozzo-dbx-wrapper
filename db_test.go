package database

import (
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	. "github.com/smartystreets/goconvey/convey"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRequestScope(t *testing.T) {
	Convey("Database", t, func() {
		mockdb, sqlmock, err := sqlmock.NewWithDSN("mock-db")
		defer mockdb.Close()
		So(err, ShouldBeNil)
		Convey("Should wrap dbx", func() {
			dbxDB, err := dbx.Open("sqlmock", "mock-db")
			So(err, ShouldBeNil)
			db := New(dbxDB)
			So(db, ShouldImplement, (*DB)(nil))
			Convey("Should begin transaction", func() {
				sqlmock.ExpectBegin()
				tx, err := db.Begin()
				So(err, ShouldBeNil)
				So(tx, ShouldImplement, (*Tx)(nil))
			})
		})
		So(sqlmock.ExpectationsWereMet(), ShouldBeNil)
	})
}
