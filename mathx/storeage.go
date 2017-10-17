package mathx

type Storeage interface {
	Init(r, c int, vals []float64)
	Get(c int) float64
	Get2(r, c int) float64

	Set(c int, val float64)
	Set2(r, c int, val float64)
}

type DataArray struct {
	rows, columns int
	data          []float64 // the contents of the matrix as one long slice.
}

func (p *DataArray) findIndex(colsAndRows ...int) int {
	var r, c int
	r = 1
	if len(colsAndRows) > 1 {
		r = colsAndRows[0]
		c = colsAndRows[1]
	} else {
		c = colsAndRows[0]
	}
	return (r-1)*p.columns + (c - 1)
}
func (p *DataArray) Init(r, c int, vals []float64) {
	if r*c != len(p.data) {
		panic("r*c != len(d.data)!")
	}
	k := 0
	valsLen := len(vals)
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			idx := p.findIndex(i, j)
			p.data[idx] = vals[k%valsLen]
			k += 1
		}
	}

}
func (p *DataArray) Set(c int, val float64) {
	idx := p.findIndex(1, c)
	p.data[idx] = val
}
func (p *DataArray) Set2(r, c int, val float64) {
	idx := p.findIndex(r, c)
	p.data[idx] = val
}

func (d *DataArray) Get(c int) float64 {
	idx := d.findIndex(1, c)
	return d.data[idx]
}
func (p *DataArray) Get2(r, c int) float64 {
	idx := p.findIndex(r, c)
	return p.data[idx]
}

func Zeros(r, c int) *DataArray {
	return &DataArray{r, c, make([]float64, r*c)}
}

func Identity(n int) *DataArray {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = 1
	}
	return A
}
