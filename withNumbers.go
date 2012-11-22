package main

import "fmt"

func main(){
	
	/* simple number manipulation */
	fmt.Println("1+1.1=",1.0+1.1)
	
	/* simple string manipulation */
	fmt.Println("length of hello world =", len("hello World"))
	fmt.Println("hello World"[1])
	fmt.Println("hello" +" world "+ "combined")
	
	/* boolean jugglery */
	fmt.Println(true&&true)
	fmt.Println(true&&false)
	fmt.Println(true||true)
	fmt.Println(true||false)
	fmt.Println(!false)
	
	/* exercise problems */
	fmt.Println(321325*424521)
}