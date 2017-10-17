package pso

import (
	. "github.com/lshengjian/pso-go/mathx"
	"sync"
	"sync/atomic"
)

type Strategy interface { //算法的移动策略
	Move(p *Particle, g, G int)
}

//Publisher
type Swarm struct {
	NP         int         // 群个体总数.
	Best       Data // 发现的最佳位置
	FEs        uint32      //函数执行总次数
	Problem    Problem
	Opt        Strategy
	sortIdxMap map[*Particle]int
	mut        sync.RWMutex   //共享数据改变锁
	DataChan  chan Data //新数据管道
	wg         sync.WaitGroup 
	IsQuick bool

}

func NewSwarm(np int, opt Strategy, p Problem) (rt *Swarm) {
	var s Swarm
	s.FEs = 0
	s.NP = np
	s.Opt = opt

	s.Problem = p
	s.DataChan = make(chan Data, np)
	return &s
}


func (s *Swarm) IncFEs() {
	atomic.AddUint32(&s.FEs, 1)
}
/*
func (s *Swarm) Finished() {
	s.wg.Done()
}*/
func (s *Swarm) GetBestX() *Vector {
	if s.IsQuick {
		s.mut.RLock()
		defer s.mut.RUnlock()
	}
	return s.Best.X.Times(1)
}
func (s *Swarm) Offset(x *Vector) *Vector {
	if s.IsQuick {
		s.mut.RLock()
	    defer s.mut.RUnlock()
	}
	return s.Best.X.Sub(x)
}
func (s *Swarm) Check(p *Particle) {
	if s.IsQuick {
	s.mut.Lock()
	defer s.mut.Unlock()
	}
	if s.Best.Value > p.Best.Value {
		s.Best.Value = p.Best.Value
		s.Best.X = p.Best.X.Times(1)
		//fmt.Println("Swarm Find Best:", p.Best.Value)

	}
}
func (s *Swarm) Init() {
	var p *Particle
	dim := s.Problem.GetDim()
	bs := s.Problem.GetBounds()
	
	s.FEs = 0
	
	s.Best.X = NewRandVector(dim, bs, 1.0)
	s.Best.Value = s.Problem.GetFunValue(s.Best.X)
	s.sortIdxMap = make(map[*Particle]int)
	for i := 0; i < s.NP; i++ {
		p = NewParticle(s)

		p.X = NewRandVector(dim, bs, 1.0)
		p.V = NewRandVector(dim, bs, 0.1)
		p.Value = s.Problem.GetFunValue(p.X)
		s.sortIdxMap[p] = i
        p.Best.X = p.X.Times(1)
		p.Best.Value = p.Value
	
		if s.Best.Value > p.Value {
			s.Best.Value = p.Value
			s.Best.X = p.X.Times(1)
		}
	}


}
func (s *Swarm) Run( G int) {
	s.Init()
	if s.IsQuick {
			for p, _ := range s.sortIdxMap {
				s.wg.Add(1)
				go p.Run(G)
			}
			s.wg.Wait()
	}else{
			for g := 0; g < G ; g++ {
				for p, _ := range s.sortIdxMap {
					p.Move(g,G)
					p.Check()
				}
			}

	}

	
	//log.Println("Time (s):", sumTime/float64(trys))
	//fmt.Println("FEs :", sumFEs/trys)
	//log.Println("pass :", sumPass*100/trys, "%")

}
