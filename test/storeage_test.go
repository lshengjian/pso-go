package mathx

import (
	"testing"
)

func Test01OneDim(t *testing.T) {
	var s Storeage
	var d, ex float64
	s = Zeros(1, 3)
	ed := []float64{0, 0, 0}
	for i := 1; i <= 3; i++ {
		d = s.Get(i)
		ex = ed[i-1]
		if d != ex {
			t.Errorf("should %f,but is %f", ex, d)
		}
	}

	ed = []float64{1, 2, 3}
	s.Init(1, 3, ed)
	for i := 1; i <= 3; i++ {
		d = s.Get(i)
		ex := ed[i-1]
		if d != ex {
			t.Errorf("should %f,but is %f", ex, d)
		}
	}
	s.Set(2, 9.0)
	d = s.Get(2)
	ex = 9.0
	if d != ex {
		t.Errorf("should %f,but is %f", ex, d)
	}

}

func Test02TwoDims(t *testing.T) {
	var s Storeage
	var d, ex float64
	s = Identity(3)
	ed := []float64{1, 1, 1}
	for i := 1; i <= 3; i++ {
		d = s.Get2(i, i)
		ex = ed[i-1]
		if d != ex {
			t.Errorf("should %f,but is %f", ex, d)
		}
	}

	ed = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s.Init(3, 3, ed)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			d = s.Get2(i, j)
			ex := ed[(i-1)*3+(j-1)]
			if d != ex {
				//t.Fail("should %f,but is %f", ex, d)
				t.Errorf("should %f,but is %f", ex, d)
			}
		}
	}
	ex = 8888.0
	s.Set2(2, 2, ex)
	d = s.Get2(2, 2)

	if d != ex {
		t.Errorf("should %f,but is %f", ex, d)
	}

}

func Benchmark01Set(b *testing.B) {
	var s Storeage
	s = Identity(3)
	for i := 0; i < b.N; i++ { //use b.N for looping
		s.Set2(1, 1, 9)
	}
}
