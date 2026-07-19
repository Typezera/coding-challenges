package main

import "fmt"

func main() {
	var num_1 int 
	var num_2 int
	var res int

	fmt.Scan(&num_1, &num_2)

	res = num_1 * num_2

	fmt.Printf("PROD = %d\n", res)

}