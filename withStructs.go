package main

import "fmt"
import "math"

type Circle struct{
	x,y,r float64
}

func area(c *Circle) float64{
	return math.Pi* c.r* c.r
}

func (c *Circle) newArea() float64{
	return math.Pi* c.r * c.r
}

func main(){

	c:= Circle{0,0,6}
//	fmt.Println(area(&c))
	fmt.Println(c.newArea())
}