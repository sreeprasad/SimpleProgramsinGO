package main

import "fmt"

func main(){

	//x :=make([]float64, 5)
	//u :=make([]float64,5, 10)
	
	/* slices using arrays */
	array := []float64{2.3,94.3,221.332,874.993,234.03874}
	m :=array[0:2]
	n :=array[0:]
	k :=array[:5]
	l :=array[:]
	
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(k)	
	fmt.Println(l)		
	
	/* appending slices */
	slice1 :=[]int{1,2,4}
	fmt.Println(slice1)
	slice2 := append(slice1,9,122)
	fmt.Println("after appending slice1",slice1)	
	fmt.Println("after appending slice2",slice2)
	
	/* coping slice2 and slice1 */
	copy(slice2,slice1)
	slice3 := make([]int,2)
	copy(slice3,slice1)
	copy(slice3,slice2)	
	fmt.Println("after copying slice1",slice1)	
	fmt.Println("after copying slice2",slice2)	
	fmt.Println("after copying slice3",slice3)		
}