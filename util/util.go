package util

import (
	"fmt"
    "io/ioutil"
    "math"

)
func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteFile(fname ,data string)  {
	err:=ioutil.WriteFile(fname,[]byte(data),0644)
    CheckError(err) 
}
type TestData struct{
    Name string
    Args string
    Data []float64
}

func (p *TestData) Min() float64{
    rt:=p.Data[0]
    for _,d:=range p.Data{
        if d < rt {
            rt=d
        }
    }
    return rt
}
func (p *TestData) Max() float64{
    rt:=p.Data[0]
    for _,d:=range p.Data{
        if d > rt {
            rt=d
        }
    }
    return rt
}
func (p *TestData) Mean() float64{
    s:=0.0
    for _,d:=range p.Data{
        s+=d
    }
    return s/float64(len(p.Data))
}
func (p *TestData) Std() float64{
    s:=0.0
    m:=p.Mean()
    for _,d:=range p.Data{
        dd:=d-m
        s+=dd*dd
    }
    return math.Sqrt(s/float64(len(p.Data)-1))
}

type ResultData struct{
    ProblemName string
 //   Tries int
    Results []*TestData
}
func (p *ResultData) SaveDataToFile(fname string)  {
    // fname:=p.ProblemName+".txt"
     str:="Method\tMean\tMin\tMax\tStd\n"
     for _,d:=range p.Results{
         str+=fmt.Sprintf("%s\t%.2G\t%.2G\t%.2G\t%.2G\n",d.Name,d.Mean(),d.Min(),d.Max(),d.Std())
     }
     WriteFile(fname,str)
}