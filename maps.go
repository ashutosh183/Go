package main


import "fmt"

func main(){
	fmt.Println("Maps")

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"UP": 266666666,
		"BIHAR": 20000000,
		"MAHARASTRA": 345278356,
		"KARNATAKA": 352683682,
	}

	statePopulations["Tamil Nadu"] = 123564818
	fmt.Println("State Populations: ", statePopulations)

	delete(statePopulations, "Tamil Nadu")

	pop, isPresent := statePopulations["Tamil Nadu"]
	fmt.Println(pop, isPresent)
	fmt.Println("State Populations: ", statePopulations)
}

