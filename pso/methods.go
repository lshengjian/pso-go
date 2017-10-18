package pso
import (
	"math/rand"
	. "github.com/lshengjian/pso-go/mathx"
)
type SPSO struct {
	
}

func (s *SPSO) Move(p *Particle, g, G int) {
	r1 := rand.Float64() * 2
	r2 := rand.Float64() * 2
	w := 0.7293//0.9 - float64(g)/float64(G)*0.5
	var d1, d2 *Vector
	d1 = p.Best.X.Sub(p.X)
	d2 = p.Swarm.OffsetOfBest(p.X)
	p.V = p.V.Times(w).Add(d1.Times(r1).Add(d2.Times(r2)))
	p.SetX(p.V.Add(p.X))

}

type BBPSO struct {
	
}

func (s *BBPSO) Move(p *Particle, g, G int) {
	b := p.Swarm.GetBestX()
	dis := b.Sub(p.Best.X).Magnitude()
	c:=b.Add(p.Best.X).Times(0.5)
	p.SetX(c.Add(NewNormalRandVector(p.Swarm.Problem.GetDim(),dis)))

}
