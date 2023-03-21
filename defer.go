package main

import (
	"fmt"
)


//Defer moves the statement to end before the function returns
func main(){
	defer fmt.Println("Defer")
	fmt.Println("start")
	defer fmt.Println("middle") //defer function allows to execute at the end
	defer fmt.Println("end")


	
}