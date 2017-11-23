package main

import fmt "fmt"

type Stringy func() string

func foo() string{
	return "Stringy function"
}

func addTwo(x int) int {
	return x + 2
}

func takesAFunction(foo Stringy){
	fmt.Printf("takesAFunction: %v\n", foo())
}

func returnsAFunction()Stringy{
	return func()string{
		fmt.Printf("Inner stringy function\n");
		return "bar" // have to return a string to be stringy
	}
}

func main(){
	takesAFunction(foo);
	var f Stringy = returnsAFunction();
	f();
	var baz Stringy = func()string{
		return "anonymous stringy\n"
	};
	fmt.Printf(baz());

	// f(x) = y
	//
	// https://imgur.com/r/linguistics/XBHub
	// λx.x
	// λx.x+2
	// λx.x+2(5)
	println(addTwo(5))
	println(func(x int) int { return x + 2 }(5))
}