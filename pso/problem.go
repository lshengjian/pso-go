package pso

import (
	. "github.com/lshengjian/pso-go/mathx"
	
)
type FitnessFunc func(v *Vector) (rt float64)

type Problem interface { //连续问题接口
	GetName() string
	GetDim() int
	GetEps() float64
	GetBound(int) (float64, float64) //获取指定维的取值范围
	GetBounds() []Bound
	GetBestValue() float64 //获取理论最优值
	GetPassValue() float64 //获取可接受值
	GetFunValue(x *Vector) float64
}
type ProblemBase struct {
	name string
	dim        int
	best, pass float64
	eps        float64
	fun        FitnessFunc
	bounds     []Bound
	
}

func NewProblemBase(name string,dim int, best, pass, eps float64, fun FitnessFunc, bounds []Bound) *ProblemBase {
	return &ProblemBase{name,dim, best, pass, eps, fun, bounds}
}
func (p *ProblemBase) GetName() string {
	return p.name
}
func (p *ProblemBase) GetDim() int {
	return p.dim
}
func (p *ProblemBase) GetEps() float64 {
	return p.eps
}
func (p *ProblemBase) GetBounds() []Bound {
	return p.bounds
}
func (p *ProblemBase) GetBound(d int) (float64, float64) {
	idx := d - 1
	if d < 1 || d > len(p.bounds) {
		idx = 0
	}
	return p.bounds[idx].From, p.bounds[idx].To
}
func (p *ProblemBase) GetBestValue() float64 {
	return p.best
}
func (p *ProblemBase) GetPassValue() float64 {
	return p.pass
}
func (p *ProblemBase) GetFunValue(x *Vector) float64 {
	return p.fun(x)
}
