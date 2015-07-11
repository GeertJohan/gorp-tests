package gorp

type Dialect interface {
	Name() string
	AddNumber(v int)
	GetSum() int
}

// WorkingDialect implements Dialect and returns correct values
type WorkingDialect struct {
	sum int
}

var _ Dialect = &WorkingDialect{}

func (w *WorkingDialect) Name() string {
	return "Working"
}

func (w *WorkingDialect) AddNumber(v int) {
	w.sum += v
}
func (w *WorkingDialect) GetSum() int {
	return w.sum
}

// BrokenDialect implements Dialect but returns faulty values to demonstrate this
type BrokenDialect struct{}

var _ Dialect = &BrokenDialect{}

func (b *BrokenDialect) Name() string {
	return "Broken"
}
func (b *BrokenDialect) AddNumber(a int) {}
func (b *BrokenDialect) GetSum() int {
	return -42
}
