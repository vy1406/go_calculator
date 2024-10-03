package prices

import (
	"fmt"
)

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPricesJob(taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		TaxRate:     taxRate,
		InputPrices: []float64{100, 200, 300},
	}
}

func (job TaxIncludedPricesJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}
