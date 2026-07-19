package main

import "fmt"

func main(){
	var grade_1 float64
	var grade_2 float64
	var res float64

	fmt.Scan(&grade_1, &grade_2)

	res = ((grade_1 * 3.5)+(grade_2 * 7.5)) / 11
	fmt.Printf("MEDIA = %.5f\n", res)

}
