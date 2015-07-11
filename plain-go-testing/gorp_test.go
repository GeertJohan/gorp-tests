package gorp

import (
	"os"
	"testing"

	. "github.com/GeertJohan/gorp-tests"
)

// global dbs var contains all *DB instances we'd like to test
var db *DB

func init() {
	// right here we can select and configure the dialect based on environment variables
	switch os.Getenv("GORP_DIALECT") {
	case "working":
		db = &DB{Dialect: &WorkingDialect{}}
	case "broken":
		fallthrough
	default:
		db = &DB{Dialect: &BrokenDialect{}}
	}
}

func TestNumbers(t *testing.T) {
	db.AddNumber(12)
	db.AddNumber(3)
	sum := db.GetSum()
	if sum != 15 {
		t.Errorf("%s: expected 15 after adding 12 and 3, got %d", db.Info(), sum)
	}
}
