package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {
	pessoa := new(Pessoa)
	pessoa.nome = "Flavio"
	pessoa.idade = 23

	fmt.Printf("Meu nome é %v e tenho %v anos de idade. ", pessoa.nome, pessoa.idade)
}
