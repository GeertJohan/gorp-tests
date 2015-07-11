package gorp

import (
	"testing"

	. "github.com/GeertJohan/gorp-tests"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type GorpSuite struct {
	DB *DB
}

var _ = Suite(&GorpSuite{
	DB: &DB{Dialect: &WorkingDialect{}},
})
var _ = Suite(&GorpSuite{
	DB: &DB{Dialect: &BrokenDialect{}},
})

func (s *GorpSuite) SetUpSuite(c *C) {
	// this doesn't even show up in the logs
	c.Logf("DB: %s", s.DB.Info())
}

func (s *GorpSuite) TestNumbers(c *C) {
	s.DB.AddNumber(12)
	s.DB.AddNumber(3)
	c.Assert(s.DB.GetSum(), Equals, 15)
}
