package pso

import (
	. "github.com/lshengjian/pso-go/mathx"

)
const G_BAD_VALUE = 1e100

type Data struct {
	X     *Vector
	Value float64
}
type Particle struct {
	X        *Vector
	V        *Vector
	Value    float64
	Best     Data
	
	Swarm    *Swarm
//	isOver   uint32
}

func NewParticle(s *Swarm) *Particle {
	var p Particle
	//dim := s.problem.GetDim()
	p.Swarm = s
	return &p

}

func (p *Particle) Check()  {
	if p.Value < p.Best.Value {
		p.Best.Value = p.Value
		p.Best.X = p.X.Times(1)
		p.Swarm.Check(p.MakeBestData())
	}
	return
}

func (p *Particle) SetX(px *Vector) {
	dim := p.Swarm.Problem.GetDim()
	for i := 1; i <= dim; i++ {
		l, h := p.Swarm.Problem.GetBound(i)
		if px.Get(i) < l {
			px.Set(i, l)
		} else if px.Get(i) > h {
			px.Set(i, h)
		}
		p.X.Set(i, px.Get(i))
	}
}
func (p *Particle) MakeBestData() Data{
	return Data{p.Best.X.Times(1),p.Best.Value}
}
func (p *Particle) Move(g, G int) {
	p.Swarm.Opt.Move(p, g, G)
	p.Value = p.Swarm.Problem.GetFunValue(p.X)
	p.Check()
/*	if p.Swarm.IsQuick{
		p.Swarm.DataChan<-p.MakeBestData()
	}
	p.Swarm.IncFEs()*/
}
/*
func (p *Particle) GetOverFlag() bool {
	return atomic.LoadUint32(&p.isOver) > 0
}*/
func (p *Particle) Run(G int) {
	go func() {
		for g := 0; g < G ; g++ {
			p.Move(g, G)
		}
		p.Swarm.wg.Done()
	}()
	
	
}
