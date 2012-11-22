package main

import "fmt"

func main(){

	/* with maps */
	x :=make(map[string] int)
	fmt.Println(x)
	x["key"]=10
	fmt.Println(x["key"])
	fmt.Println(x)
	x["value"]=20
	fmt.Println(x)
	/* after deletion */
	delete(x,"key")
	fmt.Println(x)
}