package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		// sempre tem que colocar uma condição que retorne apenas true ou false (== ou != por exemplo), não aceita if comando {}, apenas se "comando" fosse boolean.
		/* 	if comando == 1 {
			   fmt.Println("Monitorando...")
		   } else if comando == 2 {
			   fmt.Println("Exibindo logs...")
		   } else if comando == 0 {
			   fmt.Println("Saindo do programa")
		   } else {
			   fmt.Println("Não conheço este comando")
		   } */

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas" // posso declarar com := sem precisar usar var e o tipo
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int // posso declarar usando var e o tipo
	// fmt.Scanf("%d", &comando) // %d representa um inteiro - & é o endereço de memoria da variavel que quero salvar (ponteiro)
	fmt.Scan(&comandoLido) // Scan não preciso passar o tipo de dado, a variavel ja inferiu
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br/"
	sites[2] = "https://www.caelum.com.br/"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}