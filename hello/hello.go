package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const numberMonitoring = 3
const delay = 5

func main() {
	programIntroduction()

	lineSpace()

	for {

		commandMenu()

		lineSpace()

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

func lineSpace() {
	fmt.Println("--------------------------------------")
}

func commandSelector() int {
	var command int
	fmt.Scan(&command)
	return command
}

func commandMenu() {
	fmt.Println("Escolha um comando para prosseguir")
	lineSpace()
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func initMonitoring() {

	fmt.Println("Monitorando...")
	sites := archiveSitesReader()

	lineSpace()
	for i := 0; i < numberMonitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			siteTester(site)
		}
		time.Sleep(delay * time.Second)
		lineSpace()
		fmt.Println("")
	}

}

func siteTester(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

	fmt.Println("")
}

func archiveSitesReader() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(arquivo)

	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites

}
