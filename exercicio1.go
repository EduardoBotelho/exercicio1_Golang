package main


import "fmt"


var lido int64
func main() {
	fmt.Print("Entre com um número: ")
	_, err := fmt.Scan(&lido)

	if err != nil{
		fmt.Println("Erro ao ler o número",err)
		return
	}

	negativo(lido)
}

func negativo(numero int64){
	if numero > 0{
	numero = -numero
	}

	fmt.Println(numero)
}