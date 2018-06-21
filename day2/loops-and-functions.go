package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i :=0;i<10;i++{
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	
	y := x
	oldy := 0.0
	for math.Abs(y-oldy) > 0.00001{
		oldy = y
		y -= (y*y - x) / (2*y)
		fmt.Println(y)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(20))
}
