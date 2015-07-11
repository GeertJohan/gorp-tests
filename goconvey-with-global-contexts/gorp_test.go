package gorp

import (
	"testing"

	. "github.com/GeertJohan/gorp-tests"

	. "github.com/smartystreets/goconvey/convey"
)

// global dbs var contains all *DB instances we'd like to test
var dbs = make([]*DB, 0)

func init() {
	// right here we can disable and configure dialects based on environment variables
	dbs = append(dbs, &DB{Dialect: &WorkingDialect{}})
	dbs = append(dbs, &DB{Dialect: &BrokenDialect{}})
}

func TestNumbers(t *testing.T) {
	for _, db := range dbs {
		testNumbers(t, db)
	}
}

func testNumbers(t *testing.T, db *DB) {
	Convey("Given an "+db.Info()+" dialect", t, func() {
		Convey("When two values are added", func() {
			db.AddNumber(12)
			db.AddNumber(3)
			Convey("The returned value from GetSum() should be equal to their sum", func() {
				So(db.GetSum(), ShouldEqual, 15)
			})
		})
	})
}
