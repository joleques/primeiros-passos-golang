package objetos

import "fmt"

type Pessoa struct {
	Nome string
}

func (p Pessoa) Andar() {
	fmt.Println(p.Nome, "andou....")
}
