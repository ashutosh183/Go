package main

import "fmt"

type SportsPlayer struct{
	Name string
	JersyNumber int
	Ranking int
}


func main(){
	person1 := SportsPlayer{
		Name: "Virat Kohli",
		JersyNumber: 18,
		Ranking: 1,
	}

	fmt.Println(person1)

	fmt.Println(person1.JersyNumber)
}