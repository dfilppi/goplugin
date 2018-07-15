package main

import (
	"fmt"
	"os"
)

// The executable just proxies to underlying functions, which are specified in the pluging
// operation mapping or blueprint
//  arg[1] = port of context proxy ( provided by go.py )
//  arg[2] = the function to call
//  arg[3:] = the args for the function if any
func main() {

	// By convention, args[1] is a function name.  The equivalent to
	// the operation in the Python plugin

	ctx_port := os.Args[1]
	fname := os.Args[2]
	args := os.Args[3:]

	if fname == "func1" {
		func1(ctx_port, args)
	} else if fname == "func2" {
		func2(ctx_port, args)
	} else {
		fmt.Println("Function " + fname + " not found")
		os.Exit(1)
	}
}
