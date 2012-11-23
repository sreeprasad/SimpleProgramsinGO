package main

import "fmt"

func main(){

	add := func (x []int) int{
			
	 	var total int =0
	
  		for _,value:=range x	{
			fmt.Println(value)
			total+=value
		}
		return total
	}	
 
	fmt.Println(add([]int {3,2}))
}