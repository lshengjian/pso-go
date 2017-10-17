//Swarm Intelligence demo

package main

import (
	"runtime"
	"time"
	"fmt"
	"os"
	"math/rand"
	. "github.com/lshengjian/pso-go/mathx"
	. "github.com/lshengjian/pso-go/pso"
	"github.com/urfave/cli"
	"github.com/lshengjian/pso-go/util"
)
var Version = "1.0.0"
func F1(x *Vector) (rt float64) {
	d := x.Data()
	for i := 0; i < len(d); i++ {
		rt += d[i] * d[i]
	}
	return
}
// go install -ldflags "-s -w"
func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "particles, p",
			Value: 20,
			Usage: "ant poplation.",
		},
		cli.IntFlag{
			Name:  "maxIter, m",
			Value: 200,
			Usage: "try times.",
		},
		cli.BoolFlag{
			Name:  "speed, s",
			Usage: "use multi CPU cores.",
		},
		cli.StringFlag{
			Name:  "output ,o",
			Value: "./mydata.txt",
			Usage: "output file `filename`",
		},
	}
	app.Name = "PSO-GO"
	app.Authors = []cli.Author{
	    cli.Author{
			Name:  "Liu Shengjian",
			Email: "lsj178@139.com",
		},
	}
	app.Usage = "PSO demo"
	app.Version = Version
	
	app.Action = func(c *cli.Context) error {
		rand.Seed(time.Now().UnixNano())  
		bs := []Bound{Bound{-50, 50}}
		ps := [] Problem {
			NewProblemBase("X^2(d=2)",2, 0, 0.01, 1e-15, F1, bs),
			NewProblemBase("X^2(d=30)",30, 0, 1, 1e-15, F1, bs),
			NewProblemBase("X^2(d=100)",100, 0, 10, 1e-15, F1, bs),
		}
		
		pops := c.Int("particles")
		maxIter := c.Int("maxIter")
		fmt.Println("maxIter:",maxIter,"particles",pops)
		if c.Bool("speed") {
			fmt.Println("CPU cores:",runtime.NumCPU())
		}
		
		cnt:=5
		rand.Seed(time.Now().UnixNano())
		timers:=make([][]float64,len(ps))
		data:=make([]*util.TestData,len(ps))
		for i,p:=range ps{
			timers[i]=make([]float64,cnt)
			
			for t := 0; t < cnt; t++ {
				t1 := time.Now()
				s := NewSwarm(pops,  &SPSO{}, p)
				s.IsQuick=c.Bool("speed")
				s.Run(maxIter)
				
				timers[i][t]=time.Since(t1).Seconds()
				fmt.Println(t,"best:",s.Best.Value)
				//fmt.Println(t,"cost time:",timers[i][t])
			}
			data[i]=&util.TestData{p.GetName(),"",timers[i]}
		}
		
		r:=&util.ResultData{}
		for _,d :=range data	{
            r.Results=append(r.Results,d)
		}
		
		r.SaveDataToFile(c.String("output"))
		return nil
	}
	app.Run(os.Args)

}
	