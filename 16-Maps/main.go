package main

import "fmt"

/*
 * A map maps keys to values.
 * The zero value of a map is nil. A nil map has no keys, nor can keys be added.
 * The make function returns a map of the given type, initialized and ready for use.
 */

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	// 1. Basic sample
	m = make(map[string]Vertex)
	m["Kocaeli"] = Vertex{
		40.727840, 29.961800,
	}
	fmt.Println(m["Kocaeli"])

	// 2. add, delete
	cities := make(map[string]string)
	cities["IST"] = "İstanbul"
	cities["ANK"] = "Ankara"
	cities["BUR"] = "Bursa"
	fmt.Println(cities)
	fmt.Println(cities["BUR"])

	delete(cities, "ANK")
	fmt.Println(cities)

	// 3. for with map
	// * for-range is randomly iterates over map.
	cities["KOC"] = "Kocaeli"
	for key, v := range cities {
		fmt.Printf("%v : %v\n", key, v)
	}

	// 4. Get list of keys
	i := 0
	keys := make([]string, len(cities))
	for key := range cities {
		keys[i] = key
		i++
	}
	fmt.Println("Keys are", keys)

	// 5. Check key in map
	_, exists := cities["ANK"]
	if !exists {
		fmt.Printf("The %v key does not exist in map\n", "ANK")
	}

	fmt.Println("Value is", cities["ANK"])
}
