package main

import "log"

func main() {

	longStr := "hgedghdeiuiuiddiuiuieeiuiuiediuiudeiuiudeiuiued"
	shortStr := "ed"

	permCount := countPermutations(longStr, shortStr)
	log.Println("permCount: ", permCount)

}


func countPermutations(longStr, shortStr string) int {
	mapsArr := mapsInit(longStr, shortStr)
	shortStrCharMap := mapsArr[0]
	changes := mapsArr[1]
	mapChanged := false
	counter := 0

	for i := 0; i <len(longStr); i++ {
		permCheck(i, &counter, shortStrCharMap, changes, &mapChanged, longStr)
	}
	return counter
}

func permCheck(i int, counter *int,shortStrCharMap, changes map[string]int, mapChanged *bool, longStr string) {
	if val, ok := shortStrCharMap[string(longStr[i])]; ok {
		*mapChanged = true
		if val > 1 {
			shortStrCharMap[string(longStr[i])]--
			changes[string(longStr[i])]++
		}
		delete(shortStrCharMap, string(longStr[i]))

		if len(shortStrCharMap) == 0 {
			*mapChanged = false
			*counter++
			resetMap(shortStrCharMap, changes)
		}
	}else{
		if *mapChanged {
			resetMap(shortStrCharMap, changes)
		}
	}
}

func mapsInit(longStr, shortStr string) []map[string]int {
	shortStrCharMap := map[string]int{}
	changes := map[string]int{}

	for i := 0; i <len(shortStr); i++ {
		if val, ok := shortStrCharMap[string(longStr[i])]; ok {
			val++
		}
		shortStrCharMap[string(shortStr[i])] = 1
		changes[string(shortStr[i])] = 0
	}
	return []map[string]int{shortStrCharMap, changes}
}

func resetMap(myMap, changes map[string]int) {
	for k,v := range changes {
		if _, ok := myMap[k]; ok {
			myMap[k] = myMap[k] + v
		}else{
			myMap[k] = v
		}
	}
}







