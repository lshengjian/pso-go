package pso
import (

	. "github.com/lshengjian/pso-go/mathx"
)
type BBPSO struct {
	
}

func (s *BBPSO) Move(p *Particle, g, G int) {
	b := p.Swarm.GetBestX()
	dis := b.Sub(p.Best.X).Magnitude()
	c:=b.Add(p.Best.X).Times(0.5)
	p.SetX(c.Add(NewNormalRandVector(p.Swarm.Problem.GetDim(),dis)))

}
