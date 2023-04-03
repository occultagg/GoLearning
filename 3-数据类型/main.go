package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello World!"
	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.ToUpper(s1))
}
