package main

//Use next command line to run this code
//go run -gcflags "-m" main.go
//it would show that y variable is moved to head (instead of stack)
//Do not use "noinline" directive in a real code unless you have a very
//good reason to do so.

//go:noinline
func foo() int {
	x := 42
	return x
}

//go:noinline
func bar() *int {
	y := 43
	return &y
}

func main() {
	_ = foo()
	_ = bar()
}
