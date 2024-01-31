package main

import "fmt"

func main() {

	//declaração
	slice := []int{2, 4, 6, 8, 10, 70, 90}

	fmt.Printf("\nlen=%d\ncap=%d \n%v\n", len(slice), cap(slice), slice)
	//ignora tudo depois da posição 0
	fmt.Printf("\nlen=%d\ncap=%d \n%v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	//ignora tudo depois da posição 2
	fmt.Printf("\nlen=%d\ncap=%d \n%v\n", len(slice[:2]), cap(slice[:2]), slice[:2])

	//ignora tudo antes da posição 2
	fmt.Printf("\nlen=%d\ncap=%d \n%v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// aumenta tamanho do slice.
	//Atenção com listas grandes, pois dobra a capacidade
	slice = append(slice, 110)
	fmt.Printf("\nlen=%d\ncap=%d \n%v\n", len(slice), cap(slice), slice)

	slice = append(slice, 111)
	fmt.Printf("\nlen=%d\ncap=%d \n%v\n", len(slice), cap(slice), slice)

}
