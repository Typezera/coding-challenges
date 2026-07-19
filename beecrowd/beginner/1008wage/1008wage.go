package main

import "fmt"

func main(){
	var wage,overtime, tot_wage float64
	var employed int

	fmt.Scan(&employed, &wage, &overtime)

	tot_wage = wage * overtime

	fmt.Printf("NUMBER = %d\n", employed)
	fmt.Printf("SALARY = U$ %.2f\n", tot_wage)
}
