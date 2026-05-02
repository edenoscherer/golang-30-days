# 🟢 Dia 1 — Setup + CLI (OBRIGATÓRIO COMEÇAR AQUI)

## 🎯 Objetivo
Criar uma CLI simples em Go que receba parâmetros e execute lógica básica.

---

## ⚙️ Passo 1 — Instalar Go

Verificar instalação:

```bash
go version
```

Se não estiver instalado:
[https://go.dev/dl/](https://go.dev/dl/)

---

## ⚙️ Passo 2 — Criar projeto

```bash
mkdir golang-30-days
cd golang-30-days

go mod init github.com/edenoscherer/golang-30-days
```

---

## ⚙️ Passo 3 — Criar estrutura inicial

```bash
mkdir day-01-cli
cd day-01-cli

touch main.go
```

---

## ⚙️ Passo 4 — Implementar CLI

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <nome> <idade>")
		return
	}

	nome := os.Args[1]
	idade, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Idade inválida")
		return
	}

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)

	if idade >= 18 {
		fmt.Println("Status: Maior de idade")
	} else {
		fmt.Println("Status: Menor de idade")
	}
}
```

---

## ▶️ Passo 5 — Executar

```bash
go run main.go Edeno 30
```

Saída esperada:

```
Nome: Edeno
Idade: 30 anos
Status: Maior de idade
```

---

## 🧪 Exercício obrigatório

Expandir a CLI para:

* Validar entrada vazia
* Tratar erro de idade negativa
* Adicionar campo opcional: profissão

---

## 🧠 Desafio extra (recomendado)

* Transformar em binário executável:

```bash
go build -o app
./app Edeno 30
```

* Aceitar flags (ex: `--nome`, `--idade`)

---

## 📌 Entrega esperada

* Código funcionando
* Tratamento de erro básico
* Execução via CLI
* Commit realizado:

```bash
git add .
git commit -m "feat: day 01 - cli inicial"
```
