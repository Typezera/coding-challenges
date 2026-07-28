package main

import "fmt"

func main() {
	var gas, km, tot float64

	fmt.Scan(&km, &gas)

	tot = km / gas

	fmt.Printf("%.3f km/l\n", tot)
}
