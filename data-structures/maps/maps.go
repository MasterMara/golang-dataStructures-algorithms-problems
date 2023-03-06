package maps

import "fmt"

func MapsDefinitions() {

	// 1.Way
	var myMap = map[int]string{}

	myMap[1] = "mustafa"
	myMap[2] = "kadir"

	for key, value := range myMap {
		fmt.Printf("Key: %d\t Value: %s\n", key, value)
	}
	fmt.Println("************************")

	// 2.Way
	myMap2 := make(map[int]string, 3)
	fmt.Printf("Len: %v\n", len(myMap2))
	fmt.Printf("myMap Len: %v\n", len(myMap))

}
