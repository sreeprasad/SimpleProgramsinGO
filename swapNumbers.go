package main

import "fmt"

func strange(x string, y int){
	fmt.Println(x,y)
}

func swap(x,y *int){

	temp:= *x
	*x=*y
	*y=temp
	
}

func main(){
x:=2
y:=4
fmt.Println("before swap ",x,y)
swap(&x,&y)
fmt.Println("after swap ",x,y)
strange("sree",23)

}