package main

import "fmt"

func getSum(x []int) int {

	var sum int =0
	for _,value:= range x{
		sum+=value
	}
	return sum
}

func main(){
	x :=[]int{1,3,4,5,9,0,23,2}
	slice1 := x[3:5]
	fmt.Println(getSum(slice1))
}