package spectator

import "fmt"

type Measurement struct {
	id    *Id
	value float64
}

func NewMeasurement(id *Id, value float64) Measurement {
	return Measurement{id: id, value: value}
}

func (m Measurement) String() string {
	return fmt.Sprintf("M{id=%v, value=%f}", m.id, m.value)
}

func (m Measurement) Id() *Id {
	return m.id
}

func (m Measurement) Value() float64 {
	return m.value
}
