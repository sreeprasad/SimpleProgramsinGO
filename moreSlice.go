package main

import "fmt"

func main(){

	array := []int{1,2,3,4,5}
	fmt.Println(array)
	
	/* access 4 element of array */
	fmt.Println(array[4])
	
	/* adding slice */
	slice := array[:]
	fmt.Println(slice[4])

	/* length of slice */
	slice2 := make ([]int ,3, 9)
	fmt.Println(slice2)
	
	/*x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5] )*/

	/* find smallest number in list */
	x := []int{
	    48,96,86,68,
	    57,82,63,70,
	    37,34,83,27,
	    19,97, 9,17,
	}
	
	var min int =x[0]
	
	for i:=1;i<len(x);i++ {
		if min>x[i]{
			min=x[i]
		}
	} 
	fmt.Println(min)
	
	
	
}