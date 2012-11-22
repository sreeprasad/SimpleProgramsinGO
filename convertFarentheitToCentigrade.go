package main

import "fmt"

func main(){
	
	fmt.Println("Enter temperature in Farentheit ");
	
	var input float64
	
	fmt.Scanf("%f",&input)
	
	//fmt.Println(input-32)
	
	var outpu1 float64 = (((input-32)*(5))/9)
	var outpu2 float64=  (input-32) * (5/9)
	var outpu3 float64= (input -32) * 5/9
	var outpu4 float64=  ((input-32) * (5/9))	
	fmt.Println("the temperature in Centigrade is ",outpu1)
	fmt.Println("the temperature in Centigrade is ",outpu2)
	fmt.Println("the temperature in Centigrade is ",outpu3)
	fmt.Println("the temperature in Centigrade is ",outpu4)	
}