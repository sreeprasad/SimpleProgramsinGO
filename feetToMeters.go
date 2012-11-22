package main

import "fmt"

func main(){

 fmt.Println("this program converts feet to inches")
 var feet float64
 fmt.Println("enter feet ")
 fmt.Scanf("%f",&feet)
 
 var inches float64= feet*0.384
 fmt.Println("in inches ",inches)

}