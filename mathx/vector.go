package mathx

import (
	"fmt"
	"math"
	"math/rand"
)

type MathFunc func(x float64) float64
type Bound struct {
	From, To float64
}

type Vector struct {
	DataArray
}


// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

func NewVector(dim int, data []float64) *Vector {
	d := Zeros(1, dim)
	d.Init(1, dim, data)
	return &Vector{*d}
}

/*
func NewRandVector(dim int, from, to float64) *Vector {
	d := Zeros(1, dim)
	bound := to - from

	for j := 1; j <= d.columns; j++ {
		idx := d.findIndex(1, j)
		d.data[idx] = from + bound*rand.Float64()
	}

	return &Vector{*d}
}*/
func NewNormalRandVector(dim int,k float64) *Vector {
	d := Zeros(1, dim)
	for j := 1; j <= d.columns; j++ {
		idx := d.findIndex(1, j)
		d.data[idx] =rand.NormFloat64()*k
	}

	return &Vector{*d}
}

func NewRandVector(dim int, bs []Bound, k float64) *Vector {
	d := Zeros(1, dim)
	for j := 1; j <= d.columns; j++ {
		idx := d.findIndex(1, j)
		bound := (bs[0].To - bs[0].From) * k
		from := bs[0].From * k
		if j > 1 && j <= len(bs) {
			bound = (bs[j-1].To - bs[j-1].From) * k
			from = bs[j-1].From * k
		}

		d.data[idx] = from + bound*rand.Float64()
	}

	return &Vector{*d}
}
func (p *Vector) Debug() {
	fmt.Printf("%v\n", p.data[:])
	//return p.data[:]
}
func (p *Vector) Data() []float64 {
	return p.data[:]
}

func (p *Vector) Magnitude() float64 {
	sum := 0.0
	for c := 1; c <= p.columns; c++ {
		d := p.Get(c)
		sum += d * d
	}
	return math.Sqrt(sum)
}
func (p *Vector) DotProduct(B *Vector) (rt float64) {
	if len(p.data) != len(B.data) {
		panic("len(A.data) != len(B.data)!")
	}
	rt = dotProduct(p.data, B.data)
	return
}

// Add adds two matrices together and returns the resulting matrix.  To do
// this, we just add together the corresponding elements from each matrix.

func (p *Vector) Add(B *Vector) *Vector {
	if len(p.data) != len(B.data) {
		panic("len(A.data) != len(B.data)!")
	}
	C := Zeros(1, p.columns)
	for c := 1; c <= p.columns; c++ {
		C.Set(c, p.Get(c)+B.Get(c))
	}
	return &Vector{*C}
}

func (p *Vector) Sub(B *Vector) *Vector {
	if len(p.data) != len(B.data) {
		panic("len(A.data) != len(B.data)!")
	}
	C := Zeros(1, p.columns)
	for c := 1; c <= p.columns; c++ {
		C.Set(c, p.Get(c)-B.Get(c))
	}
	return &Vector{*C}
}

func (p *Vector) Times(k float64) *Vector {
	C := Zeros(1, p.columns)
	for c := 1; c <= p.columns; c++ {
		C.Set(c, p.Get(c)*k)
	}
	return &Vector{*C}
}

func (p *Vector) InitBy(f MathFunc) *Vector {
	r := Zeros(1, p.columns)
	for c := 1; c <= p.columns; c++ {
		r.Set(c, f(p.Get(c)))

	}
	return &Vector{*r}
}
