package main

import (
	"fmt"
	"primeiros-passos-golang/funcoes/objetos"
)

func main() {

	resSoma := soma(1, 2)
	fmt.Println(resSoma)

	resTotal := somaTudo(5, 6, 8, 9, 74, 25)
	fmt.Println(resTotal)

	pessoa := objetos.Pessoa{
		Nome: "Jorge",
	}

	pessoa.Andar()

}

func soma(a int, b int) int {
	return a + b
}

func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}

	return resultado
}
