package main

import "fmt"

const COMBUSTIVEL = 12

func main() {
	var horas, velocidade, total float64

	fmt.Scan(&horas, &velocidade)

	total = ((velocidade / COMBUSTIVEL) * horas)

	fmt.Printf("%.3f\n", total)
}
