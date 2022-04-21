package main

import (
	"fmt"
	"reflect"
)

// func main() {
//     var nome string = "Douglas"
//     var idade int = 24
//     var versao float32 = 1.1
//     fmt.Println("Olá sr.", nome, "sua idade é", idade)
//     fmt.Println("Este programa está na versão", versao)

//     fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
// }

/*-----Declaração curta de variáveis--------------*/

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}