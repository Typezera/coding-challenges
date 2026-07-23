package main

import "fmt"

const PI = 3.14159

func main() {
	var a, b, c, res float64

	fmt.Scan(&a, &b, &c)

	res = (a * c) / 2

	fmt.Printf("TRIANGULO: %.3f\n", res)

	res = (c * c) * PI
	fmt.Printf("CIRCULO: %.3f\n", res)

	res = ((a + b) * c) / 2
	fmt.Printf("TRAPEZIO: %.3f\n", res)

	res = b * b
	fmt.Printf("QUADRADO: %.3f\n", res)

	res = a * b
	fmt.Printf("RETANGULO: %.3f\n", res)
}
