package test

import (
	"fmt"
	"math/rand"
    "math"
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/lshengjian/pso-go/util"
)

func Test00(t *testing.T) {
	sum:=0
	for i:=0 ;i<100;i++{
	   k:= rand.NormFloat64()
	   if math.Abs(k)<0.5{
		   sum+=1
	   }
	  
	}
	fmt.Println("NormFloat64 :",sum)
	//t.pass()
}
func Test01(t *testing.T) {
	a := assert.New(t)
	demo:=&TestData{"","",[]float64{0,4,8}}
	a.Equal(0.0,demo.Min())
	a.Equal(8.0,demo.Max())
	a.Equal(4.0,demo.Mean())
	a.Equal(4.0,demo.Std())
}
func Test02(t *testing.T) {
	//a := assert.New(t)
	spso:=&TestData{"SPSO","w=0.729",[]float64{0.02,0.04,0.08}}
	bbpso:=&TestData{"BBPSO","",[]float64{0.000001,0.000002,0.0000003}}
	data:=&ResultData{"demo",[]*TestData{spso,bbpso}}
	data.SaveDataToFile()
}