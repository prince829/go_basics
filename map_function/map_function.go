package map_function

import "fmt"

func Map() {
	// A map is a collection of key-value pairs. Think of it like a dictionary
	// Keys must have the same type.
	// Values must have the same type.
	ages := make(map[string]int) //create an empty map
	ages["Alice"] = 20
	ages["John"] = 30
	fmt.Println((ages))
	people := map[string]string{
		"Alice":  "Paris",
		"John":   "London",
		"Mina":   "Tokyo",
		"Prince": "New York",
	}

	// Print all pairs
	for name, city := range people {
		fmt.Println(name, "lives in", city)
	}

	// Check if Bob exists
	city, ok := people["Bob"]
	if ok {
		fmt.Println("Bob lives in", city)
	} else {
		fmt.Println("Bob not found in the map")
	}
	delete(people, "Alice") // removes Alice from map
}
