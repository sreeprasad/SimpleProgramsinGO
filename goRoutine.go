package main

 
import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int){

	for i:=0;i<10;i++ {
		fmt.Println(n,":",i)
		amt:=time.Duration(rand.Intn(900))
		time.Sleep(time.Millisecond*amt)
	}
}

 
func main(){
	
	for i:=0;i<10;i++{
		go f(i)
	}
	fmt.Println("not waiting for go")
	var input int
	fmt.Scanln(&input)
	fmt.Println("input is ",input)
	
}