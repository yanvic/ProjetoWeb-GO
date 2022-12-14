package main

import "fmt"

func main() {
	//ContaYan := ContaConrrente{titular: "yan", numeroAgencia: 250, numeroConta: 12456, saldo: 255}
	contaYan := ContaConrrente{Titular: "yan", saldo: 300}
	contaVitor := ContaConrrente{Titular: "vitor", saldo: 200}

	status := contaYan.Tranferir(200, &contaVitor)

	fmt.Println(status)
	fmt.Println(contaYan)
	fmt.Println(contaVitor)
	//contaYan.titular = "yan"
	//contaYan.saldo = 100

	//fmt.Println(contaYan.saldo)

	//fmt.Println(contaYan.Saque(50))
	//fmt.Println(contaYan.saldo)

	//status, valor := contaYan.Depositar(500)
	//fmt.Println(status, valor)

}
