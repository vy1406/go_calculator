package prices

import (
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"` // ignore this key
	TaxRate           float64             `json:"lol_tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPricesJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{100, 200, 300},
	}
}

func (job *TaxIncludedPricesJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		texIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", texIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}
