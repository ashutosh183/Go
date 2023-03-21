package main

import "fmt"

type Animal struct{
	Name string
	Origin string
}

type Bird struct{
	Animal //go provides has a structure instead of is
	SpeedOfFly float32
	Food string
}

func main(){
	fmt.Println("Composition")

	b := Bird{}
	b.Name = "Chicken"
	b.Origin = "None"
	b.SpeedOfFly = 5.6
	b.Food = "insects"


	fmt.Println("Name: ",b.Name)
	fmt.Println(b)
}