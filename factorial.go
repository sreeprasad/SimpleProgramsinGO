package main

import "fmt"

func factorial (x uint) uint{

 if x==0{
		return 1
	}else{
		return x*(factorial(x-1))
	}
 	fmt.Println("this never executes")
	return 1
}

func main(){
	fmt.Println(factorial(5))
}