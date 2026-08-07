package main

import "fmt"

func main() {
	var valor, aux, val_100, val_50, val_20, val_10, val_05, val_02, val_01 int

	fmt.Scan(&valor)

	fmt.Println(valor)

	val_100 = valor / 100
	aux = valor % 100
	val_50 = aux / 50
	aux = aux % 50
	val_20 = aux / 20
	aux = aux % 20
	val_10 = aux / 10
	aux = aux % 10
	val_05 = aux / 5
	aux = aux % 5
	val_02 = aux / 2
	aux = aux % 2
	val_01 = aux / 1
	aux = aux % 1

	fmt.Printf("%d nota(s) de R$ 100,00\n", val_100)
	fmt.Printf("%d nota(s) de R$ 50,00\n", val_50)
	fmt.Printf("%d nota(s) de R$ 20,00\n", val_20)
	fmt.Printf("%d nota(s) de R$ 10,00\n", val_10)
	fmt.Printf("%d nota(s) de R$ 5,00\n", val_05)
	fmt.Printf("%d nota(s) de R$ 2,00\n", val_02)
	fmt.Printf("%d nota(s) de R$ 1,00\n", val_01)

}
