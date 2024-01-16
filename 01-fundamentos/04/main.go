package main

import "fmt"

type ID int

// Declaração GLOBAL
var (
	b bool = false
	c int = 10
	d string = "guipalm4"
	e float32 = 1.2
	f ID = 1
)

func main() {

	fmt.Printf("O tipo de E é %T\n", e)
	fmt.Printf("O valor de E é %v\n", e)
	fmt.Printf("O tipo de F é %T\n", f)

	// '\n' = quebra de linha

}
