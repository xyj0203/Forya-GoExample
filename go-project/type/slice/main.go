package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""
	slice := strings.Split(str, ",")
	fmt.Printf("%d", len(slice))
}
