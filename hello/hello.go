package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	programIntroduction()

	line()

	for {

		commandMenu()

		line()

		command := commandSelector()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo Logs")
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando indisponível")
			os.Exit(-1)
		}
	}
}

func programIntroduction() {
	fmt.Println("Digite o seu nome:")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão:", versao)

}

func line() {
	fmt.Println("--------------------------------------")
}

func commandSelector() int {
	var command int
	fmt.Scan(&command)
	return command
}

func commandMenu() {
	fmt.Println("Escolha um comando para prosseguir")
	line()
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func initMonitoring() {
	var site string
	fmt.Println("Insira a URL do site a ser monitorado:")
	fmt.Scan(&site)
	fmt.Println("Monitorando...")
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

	line()
}
