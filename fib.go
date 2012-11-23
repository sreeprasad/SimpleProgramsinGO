package main

import "fmt"

func fib(x int) (int){
	if x ==0{
//		fmt.Println("x ",x)
		return 0
	}
	if x==1{
//	fmt.Println("x ",x)	
		return 1
	}
//	fmt.Println("x ",x)
	return fib(x-1)+fib(x-2)
}

func main(){
	
	y:=2
	fmt.Println("fib of ",y)
	x:= fib(y)
	fmt.Println(x)
}