package deferfeatures

import "fmt"

// Let's learn defer keyword case by case from golamastery :) in golang :)

// Basic Defer Example
func BasicDeferExample() {

	// What is result ? simple: 1,2,3 yes you are right but what if use defer ? Let's see the result ?
	// My guess is 2,3,1
	// Yes answer is true but why ? --> if you use defer then compiler will execute like a last line that line
	defer fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

// Interma Defer Example
func IntermaDeferExample() {

	// What if you want to use multiple defer ?? Let's see
	// My guess is 4, 3, 2, 1 Let's see :)
	// Yes answer is true ut why ? --> if you use multiple defer in function, compiler will push in stack that lines and works like LIFO --> last in first out.

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}

// Legend Defer Example for Crystal Clear Understanding
func LegendDeferExample() {

	// Guess answer :) not hard actually. You got the idea push in stack and works like LIFO :)
	// My Guess is : 30, 2, 1, 0, 20, 10 :) Let's see
	// Great  answer is true :) you got it bro :)
	defer fmt.Println("10")
	defer fmt.Println("20")
	fmt.Println("30")
	myDeferFunc()

}

func myDeferFunc() {

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

}
