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
			Value: 200,
			Usage: "particle poplation.",
		},
		cli.IntFlag{
			Name:  "iterations, i",
			Value: 200,
			Usage: "iterations.",
		},
		cli.IntFlag{
			Name:  "tries, t",
			Value: 5,
			Usage: "tries.",
		},
		cli.BoolFlag{
			Name:  "speed, s",
			Usage: "use multi CPU cores.",
		},
		cli.StringFlag{
			Name:  "output ,o",
			Value: "p200",
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
		iterations := c.Int("iterations")
		fmt.Println("iterations:",iterations,"particles",pops)
		if c.Bool("speed") {
			fmt.Println("CPU cores:",runtime.NumCPU())
		}
		
		cnt:=c.Int("tries")
		rand.Seed(time.Now().UnixNano())
		timers:=make([][]float64,len(ps))
		datas:=make([][]float64,len(ps))
		timeData:=make([]*util.TestData,len(ps))
		valueData:=make([]*util.TestData,len(ps))
		for i,p:=range ps{
			timers[i]=make([]float64,cnt)
			datas[i]=make([]float64,cnt)
			for t := 0; t < cnt; t++ {
				t1 := time.Now()
				s := NewSwarm(pops,  &SPSO{}, p)
				s.IsQuick=c.Bool("speed")
				s.Run(iterations)
				
				timers[i][t]=time.Since(t1).Seconds()
				fmt.Println(t,"best:",s.Best.Value)
				datas[i][t]=s.Best.Value
				//fmt.Println(t,"cost time:",timers[i][t])
			}
			timeData[i]=&util.TestData{p.GetName(),"",timers[i]}
			valueData[i]=&util.TestData{p.GetName(),"",datas[i]}
		}
		
		r1:=&util.ResultData{}
		r2:=&util.ResultData{}
		for _,d :=range timeData	{
            r1.Results=append(r1.Results,d)
		}
		for _,d :=range valueData	{
            r2.Results=append(r2.Results,d)
		}
		flag:=""
		if c.Bool("speed"){
			flag="Q"
		}
		r1.SaveDataToFile("T-"+c.String("output")+flag+".txt")
		r2.SaveDataToFile("V-"+c.String("output")+flag+".txt")
		return nil
	}
	app.Run(os.Args)

}
	