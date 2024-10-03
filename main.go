package main

import (
	"example.com/calculator/cmdmanager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		priceJob := prices.NewTaxIncludedPricesJob(cmdm, taxRate)
		priceJob.Process()
	}

}
