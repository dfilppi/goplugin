package main

import "fmt"

// Example plugin function
func func1(ctx_port string, args []string) {
	fmt.Println("FUNC1 CALLED")
	res, err := call_context(ctx_port, "node", "properties", "prop1")
	fmt.Printf("GOT VALUE of PROP1 = %s\n", res)
	call_context(ctx_port, "logger", "info", "LOGGING SOMETHING FROM PLUGIN")
}

// Example plugin function
func func2(ctx_port string, args []string) {
	fmt.Println("FUNC2 CALLED")
}
