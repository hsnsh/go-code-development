package main

import (
	"fmt"
	"github.com/hsnsh/go-code-development/calc"
)

func main() {
	fmt.Println("Code Development Sample")

	cal := calc.Calculator{}
	fmt.Println(cal.Add(1, 1))
}
