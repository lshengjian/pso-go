package mathx

import (
	"testing"
)

func Test01Matrix(t *testing.T) {

	m1 := NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
	e := m1.Get2(1, 1)
	if e != 1 {
		t.Errorf("m[1,1] should 1,but is %f", e)
	}
	m1.Set2(1, 1, 9)
	e = m1.Get2(1, 1)
	if e != 9 {
		t.Errorf("m[1,1] should 9,but is %f", e)
	}
	m2 := m1.Times(2.0)
	e = m2.Get2(2, 3)
	if e != 12 {
		t.Errorf("m[2,3] should 12,but is %f", e)
	}
}
