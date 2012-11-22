package main

import "fmt"

func main (){

	var x [5] float64
		 x[0] = 98
	     x[1] = 93
	     x[2] = 77
	     x[3] = 82
	     x[4] = 83
	
	var total float64=0.0
	var totalStudents float64 = 5
	
  	for _,value := range x{
		total+=value
	}	

	fmt.Println("total is ",total)
	fmt.Println("number of students ",totalStudents)
	var average float64 = total/totalStudents
	fmt.Println("average is ", average)
	
	
}