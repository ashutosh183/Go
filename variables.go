package main

import "fmt"



func main(){
	//variable declaration
	var i int
	i = 42
	fmt.Println(i)


	//go compiler automatically identifies the type
	j := 56
	fmt.Println(j)

	//format printing
	fmt.Printf("J value is %v and it's type is %T", j, j)
	fmt.Println()

	name := "Ashutosh Pandey"
	fmt.Printf("Name is %s", name)

	fmt.Println()
	price := 87.9
	fmt.Printf("Price is %f", price)

	
}