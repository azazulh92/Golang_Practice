package main

import "fmt"

// replace a character of a string

func main() {
	var s string = "Azazul"
	repl := s[:1] + "j" + s[2:]
	fmt.Printf("In string Azazul , z replace as j:=%s", repl)
}
