package main

import (
	"errors"
	"fmt"
)

func main() {

	println("Soma 50 + 10: ", sum(50, 10))
	println("Soma 50 + 10 com shorthand: ", sumShortHand(50, 10))
	sum, isGraterThan50 := sumIsGraterThan50(50, 10)
	println("Soma 50 + 10 com retorno duplo(resultado,bool): ", sum, isGraterThan50)

	sumLessThan50, err := sumIfLessThan50(50, 10)

	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println(sumLessThan50)
	}

}

func sum(a int, b int) int {
	return a + b
}

func sumShortHand(a, b int) int {
	return a + b
}

// retorno com mais de um valor
func sumIsGraterThan50(a, b int) (int, bool) {
	sum := a + b
	if sum > 50 {
		return sum, true
	}
	return sum, false
}

func sumIfLessThan50(a, b int) (int, error) {
	sum := a + b
	if sum > 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return sum, nil
}
