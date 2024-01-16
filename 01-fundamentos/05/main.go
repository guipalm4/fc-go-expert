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

	var array [3] int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	fmt.Println(len(array))
	//ultimo elemento
	fmt.Println(array[len(array) - 1])

	for i, v := range array {
		fmt.Printf("O valor do indice %v é %v\n", i, v)
	}
}
