package main

import "fmt"

var globalString string ="my name is Anthony Golzalviz"
func main(){
	
	/* enclosing multiple variables to single block */
	var (
		
		x1 ="vaishnavi"
		x2= "BlackRock"
		x3 = 34.22
		x4 = 98
		x5 = true
	
	)
	
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
				
	var x string ="hello Sreeprasad"
	fmt.Println(x)
	x="hello Govindankutty"
	fmt.Println(x)
	
	x="sree"
	var y string ="sree"
	fmt.Println(x==y)
	
	x="sree"
	y="govindankutty"
	fmt.Println(x==y)
	
	u:=34.87
	fmt.Println(u)
	
	fmt.Println(globalString)
}

func x(){
	fmt.Println(globalString)
}