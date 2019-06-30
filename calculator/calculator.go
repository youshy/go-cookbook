package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(x float64, y float64) float64 {
	return x + y
}

func substract(x float64, y float64) float64 {
	return x - y
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func divide(x float64, y float64) float64 {
	return x / y
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Provide one or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	x, _ := strconv.ParseFloat(arguments[1], 64)
	y, _ := strconv.ParseFloat(arguments[2], 64)
	operator := arguments[3]

	if operator == "add" {
		fmt.Println(add(x, y))
	} else if operator == "substract" {
		fmt.Println(substract(x, y))
	} else if operator == "multiply" {
		fmt.Println(multiply(x, y))
	} else if operator == "divide" {
		fmt.Println(divide(x, y))
	}
}
