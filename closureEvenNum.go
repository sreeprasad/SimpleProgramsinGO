package main

import "fmt"

func nextEvenGen() func() uint {
	i:= uint(0)
 	return func() (ret uint)  {
		ret =i
		i+=2
 		return 
	}
}

 

func main() {
 
	 ne := nextEvenGen()
	 fmt.Println(ne())
	 fmt.Println(ne())
	 fmt.Println(ne())		
	
}



/*func main(){

	
	next:=nextEven()
	fmt.Println(next)
	fmt.Println(next)
	
}*/