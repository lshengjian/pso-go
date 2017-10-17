package mathx

// dotProduct calculates the algebraic dot product of two slices.  This is just
// the sum  of the products of corresponding elements in the slices.  We use
// this when we multiply matrices together.

func dotProduct(a, b []float64) float64 {
	var total float64
	for i := 0; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total
}

type Matrix struct {
	DataArray
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

func NewMatrix(r, c int, data []float64) *Matrix {
	d := Zeros(r, c)
	d.Init(r, c, data)
	return &Matrix{*d}
}

// Column returns a slice that represents a column from the matrix.
// This works by examining each row, and adding the nth element of
// each to the column slice.

func (p *Matrix) Column(n int) []float64 {
	col := make([]float64, p.rows)
	for i := 1; i <= p.rows; i++ {
		col[i-1] = p.Row(i)[n-1]
	}
	return col
}

// Row returns a slice that represents a row from the matrix.

func (p *Matrix) Row(n int) []float64 {
	return p.data[p.findIndex(n, 1):p.findIndex(n, p.columns+1)]
}

func (p *Matrix) Multiply(B *Matrix) *Matrix {
	C := Zeros(p.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		A_row := p.Row(r)
		for c := 1; c <= C.columns; c++ {
			B_col := B.Column(c)
			C.Set2(r, c, dotProduct(A_row, B_col))
		}
	}
	return &Matrix{*C}
}

// Add adds two matrices together and returns the resulting matrix.  To do
// this, we just add together the corresponding elements from each matrix.

func (p *Matrix) Add(B *Matrix) *Matrix {
	C := Zeros(p.rows, p.columns)
	for r := 1; r <= p.rows; r++ {
		for c := 1; c <= p.columns; c++ {
			C.Set2(r, c, p.Get2(r, c)+B.Get2(r, c))
		}
	}
	return &Matrix{*C}
}
func (p *Matrix) Times(k float64) *Matrix {
	C := Zeros(p.rows, p.columns)
	for r := 1; r <= p.rows; r++ {
		for c := 1; c <= p.columns; c++ {
			C.Set2(r, c, p.Get2(r, c)*k)
		}
	}
	return &Matrix{*C}
}
