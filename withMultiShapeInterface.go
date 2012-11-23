package main

import "fmt"
import "math"

type Shapes interface {
	area() float64 	
}

type MultipleShapes struct {
	shapes []Shapes
}

type Circle struct{
	x,y,r float64
}

type Rectangle struct {
	l,b float64
}

func (m *MultipleShapes) area() float64{
	totalarea :=0.0
	for _,shape:= range m.shapes{
		totalarea+=shape.area()
	}
	return totalarea;
}

func (r Rectangle) area() float64{
	return r.l*r.b
}

func (c Circle) area() float64{
	return math.Pi * c.r * c.r
}

func main(){
	c :=Circle{0,0,8}
	r :=Rectangle{3,9}
	fmt.Println( c.area())
	fmt.Println( r.area())	
	fmt.Println(area())
	
}