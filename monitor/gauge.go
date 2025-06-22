package monitor

type Value float64
type Unit string

type SimpleValue struct {
	Current Value
	Unit    Unit
}

type Gauge struct {
	Min, Max, Current Value
	Unit              Unit
}
