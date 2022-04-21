package main

import "fmt"



func main() {
    nome := "Lúcio"
    versao := 1.1
    fmt.Println("Olá sr.", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)

    if comando == 1 {
        fmt.Println("Monitorando...")
    } else if comando == 2 {
        fmt.Println("Exibindo Logs...")
    } else if comando == 0 {
        fmt.Println("Saindo do programa")
    } else {
        fmt.Println("Não conheço este comando")
    }
/*---------------Outra forma de fazer----------------*/

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 3:
        fmt.Println("Saindo do programa")
    default:
        fmt.Println("Não conheço este comando")
    }
    // o switch não precisa utilizar o comando break para parar a execução....ele por si só já vai sair sa função caso alguem escolher um item



}