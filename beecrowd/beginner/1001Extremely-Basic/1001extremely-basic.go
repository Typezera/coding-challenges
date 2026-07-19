package main

import (
	"fmt"
)

func main(){
	var val_1 int
	var val_2 int
	var res int
	fmt.Scan(&val_1, &val_2)

	res = val_1 + val_2

	fmt.Println("X =", res)
}