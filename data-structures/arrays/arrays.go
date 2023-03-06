package arrays

import "fmt"

func ArrayDefinitions() {

	// 1.Way
	var ages = [3]int{1, 2, 3}
	var numbers [3]int

	// 2.Way
	cities := [2]string{"Istanbul", "Paris"}

	// 3.Way
	names := make([]string, 3)

	// 4.Way
	myMatrix := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	for _, age := range ages {
		fmt.Printf("%d\t", age)
	}
	fmt.Printf("\n")

	for _, city := range cities {
		fmt.Printf("%s\t", city)
	}
	fmt.Printf("\n")

	for _, number := range numbers {
		fmt.Printf("%d\t", number)
	}
	fmt.Printf("\n")

	for _, name := range names {
		fmt.Printf("%s\t", name)
	}
	fmt.Printf("\n")

	for _, row := range myMatrix {
		for _, item := range row {
			fmt.Printf("%d\t", item)
		}
	}

}

/* Some Properties

1-)Arrays is fixed size.
2-)


*/
