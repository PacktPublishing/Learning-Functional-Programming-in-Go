package main

func addTwo(x int) int {
	return x + 2
}

func main() {
	println(addTwo(5))  // named function

	println(func(x int) int {return x + 2}(5)) // anonymous function

	val := func(x int) int {return x + 2}(5) // function expression
	println(val)
}

