package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const numberMonitoring = 3
const delay = 10

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
	// fmt.Println("Insira a URL do site a ser monitorado:")
	// fmt.Scan(&site)
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br", "https://www.globo.com"}
	fmt.Println(sites)

	line()
	for i := 0; i < numberMonitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			siteTester(site)
		}
		time.Sleep(delay * time.Second)
		line()
		fmt.Println("")
	}

}

func siteTester(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

	fmt.Println("")
}
