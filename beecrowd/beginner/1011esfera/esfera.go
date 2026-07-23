package main

import "fmt"

const PI = 3.14159

func main() {
	var vol, r float64

	fmt.Scan(&r)

	vol = ((4.0 / 3.0) * PI * (r * r * r))

	fmt.Printf("VOLUME = %.3f\n", vol)
}
