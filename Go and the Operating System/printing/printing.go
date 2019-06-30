package main

import (
	"fmt"
)

func main() {
	v1 := "123"
	v2 := 123
	v3 := "GO GET ME!\n"
	v4 := "abcdefg"

	// Prints without newline at the end
	fmt.Print(v1, v2, v3, v4)
	// Prints with newline at the end
	fmt.Println()
	fmt.Println(v1, v2, v3, v4)
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	// Prints only if the format specifier is defined for each variable
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)
}
