package main

import "fmt"

func main(){

	     xs := []float64{98,93,77,82,83}
    	fmt.Println(average(xs))
		x,y,z := f()
		fmt.Println(x, " and ", y, " and ",z)
		fmt.Println("with new average ",newAverage(23.33,12.397,93.93,23.8989,831.123))
		
}

func newAverage(x ...float64) float64{
	var total float64 =0.0
	
	for i:=0;i<len(x);i++{
		total+=x[i]
	}
	
	return total/5.0
}

func f() ( int, float64, string){
	return 4,5.823797,"sreeprasad"
}

func average(x []float64 ) float64{

	var	total float64 = 0.0
	for i:=0;i<len(x);i++ {
		total+=x[i]
	} 
	var average float64 = total/float64(5)
	return average

}