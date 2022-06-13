package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0 ; i <= 100 ; i++{
		z -= (z*z -x) / (2*x)  
	}
	return z
}

func main (){
	fmt.Println(Sqrt(2),math.Sqrt(2))
	fmt.Println(Sqrt(5),math.Sqrt(5))
}