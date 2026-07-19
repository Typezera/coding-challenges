package main

import "fmt"

const PI = 3.14159

func main(){
	var area float64
	var res float64

	fmt.Scan(&area)

	res = (area*area) * PI

	fmt.Printf("A=%.4f\n", res)
}