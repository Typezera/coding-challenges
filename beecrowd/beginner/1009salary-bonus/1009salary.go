package main

import "fmt"

const BONUS = 0.15

func main() {
	var name string
	var salary, sales, total float64

	fmt.Scan(&name, &salary, &sales)

	total = (sales * BONUS) + salary

	fmt.Printf("TOTAL = R$ %.2f\n", total)

}
