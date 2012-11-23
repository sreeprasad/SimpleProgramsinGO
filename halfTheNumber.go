package main

import "fmt"

func getReturn(x int) (int,bool){

 	fmt.Println("x is ",x)
	y:= x/2
	if y%2 ==0{
		return y,true
		}
	return y,false

}

func main(){

 
	fmt.Println(getReturn(8))

}