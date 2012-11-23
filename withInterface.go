package main

import "fmt"
import "math"

type Circle struct{
	x,y,r float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
 	len, bre float64
}

func (r Rectangle) perimeter() float64{
	return 2*(r.len+r.bre)
}
func (c Circle) perimeter() float64{
	return 2*math.Pi*c.r
}
func (r Rectangle) area() float64{
	return r.len * r.bre
}

func (c Circle)area() float64{
	return math.Pi* c.r * c.r
}

func totalPerimeter(shapes ...Shape) float64{
	peri :=0.0 
	for _,shape:= range shapes{
		peri+=shape.perimeter()
	}
	return peri
}

func totalarea(shapes ...Shape) float64{

	area:=0.0
	
	for _,shape :=range shapes{
		area+=shape.area()
	}
	
	return area;
	
}

func main(){

	c :=Circle{0,0,90}
	r :=Rectangle{890.922,40.34234}
	fmt.Println(totalarea(&c,&r))
	fmt.Println(totalPerimeter(&c,&r))
}