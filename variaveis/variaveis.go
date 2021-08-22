package main

import (
	"fmt"
	"primeiros-passos-golang/variaveis/operacoes"
)

func main() {

	a := 20
	b := "Teste"
	c := 3.144
	d := false
	e := `teste teste
	
	Legal
	`

	f := operacoes.Soma(1, 2)

	g := operacoes.Soma10(12)

	h := operacoes.A

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)

}
