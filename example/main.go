package main

import (
	"fmt"

	"github.com/kompiuter/equal"
)

func main() {
	slice := []string{"hello", "world", "gopher"}
	str := "gopher"
	fmt.Println(equal.StringContains(slice, str))
}
