package main

import "fmt"

func oddGenerator() func() uint{
	
	x:=uint(1)
	return func() (ret uint){
		
		ret=x;
		x+=2;
		return
	}
}

func main(){
 	x:= oddGenerator()
	fmt.Println(x()) // 1
	fmt.Println(x()) // 3
	fmt.Println(x()) // 5	
}