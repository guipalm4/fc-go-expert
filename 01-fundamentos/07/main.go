package main

import "fmt"

func main() {

	//declaração
	salarios := map[string]int{"Guilherme": 15800, "Victoria": 4000}
	salarios2 := make(map[string]int)
	salarios2["Guilherme"] = 16000

	//acesso
	fmt.Println(salarios["Guilherme"])
	fmt.Println(salarios)

	//exclusão
	delete(salarios, "Victoria")
	fmt.Println(salarios)

	//atribuição
	salarios["Guilherme"] = 1600
	fmt.Println(salarios["Guilherme"])

	salarios["Victória"] = 4000

	//percorrendo usando as variaveis de chave e valor
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é R$ %d!\n", nome, salario)
	}

	//percorrendo usando apenas o valor
	for _, salario := range salarios {
		fmt.Printf("O salário é R$ %d!\n", salario)
	}

}
