package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) TrocaNome(nomeNovo string) {
	c.Nome = nomeNovo
	fmt.Println("Nome Trocado sem Ponteiro:", c.Nome)
}

func (c *Carro) TrocaNomePonteiro(nomeNovo string) {
	c.Nome = nomeNovo
	fmt.Println("Nome Trocado com Ponteiro:", c.Nome)
}

func main() {

	car := Carro{
		Nome: "Clio",
	}

	car.TrocaNome("Renegade")
	fmt.Println("Nome =====> ", car.Nome)

	car.TrocaNomePonteiro("Duster")
	fmt.Println("Nome =====> ", car.Nome)
}
