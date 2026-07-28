package main

import (
	"fmt"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b, c, maiorAB, maior int

	fmt.Scan(&a, &b, &c)

	maiorAB = (a + b + absInt(a-b)) / 2
	maior = (maiorAB + c + absInt(maiorAB-c)) / 2

	fmt.Printf("%d eh o maior\n", maior)
}
