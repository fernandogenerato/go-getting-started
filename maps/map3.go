package main

import "fmt"

func main() {
	type State struct {
		name       string
		population int
		capital    string
	}

	//make sp object
	sp := State{
		name:       "São Paulo",
		population: 12441414,
		capital:    "SP",
	}

	// create map with size 3
	states := make(map[string]State, 3)
	states["SP"] = sp
	states["RJ"] = State{
		name:       "Rio de Janeiro",
		population: 1234121,
		capital:    "RJ",
	}

	//lookup usage to verify if contains "GO"
	state, found := states["GO"]

	if found {
		fmt.Println(state)
	}

	//delete SP from map
	delete(states, "SP")
	fmt.Println(states)

	//add SP
	states["SP"] = sp

	//interate map
	for key, value := range states {
		fmt.Printf("population : %d,  capital : %s, population : %s\n",
			value.population,
			key,
			value.name)

	}
}
