package main

import "fmt"

// Example plugin function
func func1(ctx_port string, args []string) {
	fmt.Println("FUNC1 CALLED")
	call_context(ctx_port, "node", "name")
	call_context(ctx_port, "logger", "info", "LOGGING SOMETHING FROM PLUGIN")
}

// Example plugin function
func func2(ctx_port string, args []string) {
	fmt.Println("FUNC2 CALLED")
}
