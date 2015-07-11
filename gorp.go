package gorp

import "fmt"

type DB struct {
	Dialect Dialect
}

func (d *DB) Info() string {
	return fmt.Sprintf("dialect:%s", d.Dialect.Name())
}

func (d *DB) AddNumber(v int) {
	d.Dialect.AddNumber(v)
}

func (d *DB) GetSum() int {
	return d.Dialect.GetSum()
}
