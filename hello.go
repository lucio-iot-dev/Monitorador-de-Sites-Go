package main

import (
	"fmt"
	"net/http"
	"os"
)



func main() {

    exibeIntroducao()

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

    fmt.Println("O valor da variável comando é:", comandoLido)

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)
    
    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}