package mathx

import (
	"math"
	"testing"
)

func Test01Vector(t *testing.T) {
	m1 := NewVector(2, []float64{1, 2})
	e := m1.Get(2)
	if e != 2 {
		t.Errorf("m[2] should 2,but is %f", e)
	}
	m2 := m1.Times(2.0)
	e = m2.Get(2)
	if e != 4 {
		t.Errorf("m[2] should 4,but is %f", e)
	}

	m1.Set(2, 9)
	e = m1.Get(2)
	if e != 9 {
		t.Errorf("m[2] should 9,but is %f", e)
	}

}

func Test02Vector(t *testing.T) {
	m1 := NewVector(2, []float64{1, 2})
	m2 := m1.InitBy(math.Sin)
	ex := math.Sin(2)
	d := m2.Get(2)

	if d != ex {
		t.Errorf("should %f,but is %f", ex, d)
	}

}
