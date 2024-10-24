package main

import "flag"

func main() {
	var name = getTheFlag()
	flag.Parse()
	println("Hello, %v!\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "the greeting object")
}
