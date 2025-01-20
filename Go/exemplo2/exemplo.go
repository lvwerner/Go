package main

import "fmt"

func main() {
	fmt.Println("Digite o seu nome:")
	var nome string        //declara a variavel nome
	fmt.Scanf("%s", &nome) // Armazena o nome do usuario na variavel nome
	fmt.Printf("Seja bem-vindo, %s!", nome)
}
