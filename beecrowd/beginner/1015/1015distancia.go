package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, res1, res2, tot float64

	fmt.Scan(&x1, &y1, &x2, &y2)

	res1 = math.Pow((x2 - x1), 2)
	res2 = math.Pow((y2 - y1), 2)

	tot = math.Sqrt(res1 + res2)

	fmt.Printf("%.4f\n", tot)
}
