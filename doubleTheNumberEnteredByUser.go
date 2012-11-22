package main

import "fmt"

func main(){

	fmt.Println("enter number")
	var input float64
	fmt.Scanf("%f",&input)
	
	output := input*2
	
	fmt.Println("the double of your input is ",output)
}