package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome  string
	Idade int `json:"idade"`
}

type Vendedor struct {
	Pessoa
	Codigo int
}

func (pessoa Pessoa) Andar() {
	fmt.Println(pessoa.Nome, "andou.....")
}

func main() {
	jorge := Pessoa{
		Nome:  "Jorge Oleques",
		Idade: 42,
	}
	jorge.Andar()

	jacque := Pessoa{"Jacqueline", 40}
	jacque.Andar()

	paulo := Vendedor{
		Pessoa: Pessoa{
			Nome:  "Paulo Roberto",
			Idade: 56,
		},
		Codigo: 565659898,
	}

	paulo.Andar()

	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", jorge.Nome, jorge.Idade)
	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", jacque.Nome, jacque.Idade)
	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", paulo.Nome, paulo.Idade)

	//Ponteiros....

	carlos := Pessoa{
		Nome:  "Carlos Roberto",
		Idade: 78,
	}
	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", carlos.Nome, carlos.Idade)

	// sem passar por referencia
	alterarNomePessoaSemReferencia(carlos, "João Carlos")
	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", carlos.Nome, carlos.Idade)

	// passando por referencia
	alterarNomePessoaPorReferencia(&carlos, "João Carlos")
	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", carlos.Nome, carlos.Idade)

	// tratando json
	jose := Pessoa{
		Nome:  "José Carlos",
		Idade: 78,
	}
	joseJson, err := json.Marshal(jose)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(joseJson))

	joseCarlos := Pessoa{}
	joseCarlosJson := `{"nome":"José Carlos","idade":78}`
	json.Unmarshal([]byte(joseCarlosJson), &joseCarlos)

	fmt.Printf("Usuario %s com idade %d foi cadastrado. \n", joseCarlos.Nome, joseCarlos.Idade)

}

func alterarNomePessoaSemReferencia(pessoa Pessoa, nomeNovo string) {
	pessoa.Nome = nomeNovo
}
func alterarNomePessoaPorReferencia(pessoa *Pessoa, nomeNovo string) {
	pessoa.Nome = nomeNovo
}
