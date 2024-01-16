package main

import "fmt"

const a = "Hello, world!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Guilherme"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	// \n = new line	%T = tipo %v = valor
	fmt.Printf("O tipo de f é %T\n", f)
	fmt.Printf("O valor de f é %v", f)
}

