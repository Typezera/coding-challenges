package main

import "fmt"

func main() {
	var priceFirstParts, priceSecondParts, totPay float64
	var codePartsOne, firstTotParts, codeSecondParts, secondTotParts float64

	fmt.Scan(&codePartsOne, &firstTotParts, &priceFirstParts)
	fmt.Scan(&codeSecondParts, &secondTotParts, &priceSecondParts)

	totPay = ((priceFirstParts * firstTotParts) + (priceSecondParts * secondTotParts))

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totPay)
}
