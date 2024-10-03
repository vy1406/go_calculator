package main

import "example.com/calculator/prices"

func main() {
	taxRates := []float64{0, 0.7, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPricesJob(taxRate)
		priceJob.Process()
	}

}
