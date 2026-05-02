package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 { // Verifica se os argumentos foram fornecidos
		fmt.Println("Uso: go run main.go <nome> <idade> <profissão>") // Exibe uma mensagem de uso e sai do programa
		return
	}

	nome := os.Args[1]                     // Acessa o primeiro argumento (nome) e armazena em uma variável
	idade, err := strconv.Atoi(os.Args[2]) // Converte o segundo argumento (idade) de string para inteiro e armazena em uma variável
	work := os.Args[3]

	if err != nil { // Verifica se houve um erro na conversão da idade
		fmt.Println("Idade inválida") // Exibe uma mensagem de erro caso a idade não seja um número válido e sai do programa
		return
	}
	if idade < 0 {
		fmt.Println("A Idade não pode ser menor que 0")
		return
	}

	fmt.Printf("Nome: %s\n", nome)        //	Exibe o nome fornecido pelo usuário
	fmt.Printf("Idade: %d anos\n", idade) // Exibe a idade fornecida pelo usuário
	fmt.Printf("Profissão: %s\n", work)   // Exime a profissão fornecida pelo usuário

	if idade >= 18 {
		fmt.Println("Status: Maior de idade")
	} else {
		fmt.Println("Status: Menor de idade")
	}
}
