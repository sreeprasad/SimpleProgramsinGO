package main

import "fmt"
func change(x *int) {
	*x=0
 
}

func main(){
	
	x:=7
	fmt.Println(x)
	change(&x)
	fmt.Println(x)
}