package main

import "fmt"

func main() {
	var first_val int
	var second_val int 
	var sum int 

	fmt.Scan(&first_val, &second_val)

	sum = first_val + second_val

	fmt.Printf("SOMA = %d\n", sum)

}
