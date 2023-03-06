package slices

import "fmt"

func SliceDefinitions() {

	// 1.Way
	var mySlice []int

	// 2.Way
	var mySlice2 = []int{}

	// 3.Way
	var myArray = [6]int{1, 2, 3, 4, 5, 6}
	mySlice3 := myArray[2:5]

	// 4. Way
	mySlice4 := make([]int, 2, 3) //len, cap

	// 5. Way
	mySlice5 := make([]int, 3) //len, cap is equal

	fmt.Printf("mySlice Cap:%d\t Len:%d\n", cap(mySlice), len(mySlice))
	fmt.Printf("mySlice2 Cap:%d\t Len:%d\n", cap(mySlice2), len(mySlice2))
	fmt.Printf("mySlice3 Cap:%d\t Len:%d\n", cap(mySlice3), len(mySlice3))
	fmt.Printf("mySlice4 Cap:%d\t Len:%d\n", cap(mySlice4), len(mySlice4))
	fmt.Printf("mySlice5 Cap:%d\t Len:%d\n", cap(mySlice5), len(mySlice5))

}

/* Slice Properties

1-)the length of a slice can grow and shrink as you see fit.

*/
