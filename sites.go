package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	// "io/ioutil"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 1

func main() {

	exibeIntroducao()
	leSitesDoArquivo()

	for {

		exibirMenu()

		comando := lerComando()
		/*---- no if sempre retornará booleano-----*/
		// if comando == 1 {
		//     fmt.Println("Monitorando...")
		// } else if comando == 2 {
		//     fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		//     fmt.Println("Saindo do programa")
		// } else {
		//     fmt.Println("Não conheço este comando")
		// }
		/*---------------Outra forma de fazer----------------*/

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
		// o switch não precisa utilizar o comando break para parar a execução....ele por si só já vai sair sa função caso alguem escolher um item
	}
}

func exibeIntroducao() {
	nome := "Lúcio"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

    sites := leSitesDoArquivo()

	// esse metodo pode ser subtituido em go por range
	// for i := 0; i < len(sites); i ++ {
	// fmt.Println(sites[i])
	// }

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	// arquivo, err := os.Open("sites.txt") --->retorna um ponteiro ,ou seja um numero do endereço de memoria
	/*----com isso podemos usar o comando abaixo-----*/

	// arquivo, err := ioutil.ReadFile("sites.txt")  // ler tudo de um aso vez. retorna um slice de bytes OBS: ABAIXO NO FMT É SO PEDIR PARA CONVERTER EM STRING
	//TEMOS UM PROBLEMA ESTA FUNÇÃO IOUTIL NÃO LÊ LINHA A LINHA ELE LE O TEXT TODO DE UMA VEZ E PRECISAMOS LER LINHA A LINHA
	/*por isso vamos implementar o seguinte comando:*/

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo) // ler linha a linha

	// leitor.ReadString('\n') // '\n' quer dizer que vai ler ate na quebra de linha
	// tem que ser aspas simples para ele não entender que é uma string e sim uma quebra de linha

	for {

		linha, err := leitor.ReadString('\n')  //Ler ate a quebra de linha
		linha = strings.TrimSpace(linha) // esta funcionalidade é não dar quebra de linha.nada após a string e válido
		
        sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

    arquivo.Close()
       
    return sites
}
