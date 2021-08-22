package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	log.Output(0, "Inicio do main.... ")

	testeComFuncaoErro(5, 2)
	testeComFuncaoErro(5, 10)
	log.Output(0, "Fim do main.... ")
}

func testeComFuncaoErro(a int, b int) {
	soma1, error1 := soma(a, b)

	log.Output(0, "teste de log.... ")
	if error1 != nil {
		log.Output(0, "Error.... ")
		log.Fatalln(error1.Error())
	} else {
		log.Output(0, "Success.... ")
		fmt.Println(soma1)
	}
}

func soma(a int, b int) (int, error) {
	result := a + b

	if result > 10 {
		return 0, errors.New("resultado maior que 10")
	}

	return result, nil
}
