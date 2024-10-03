package prices

import (
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/filemanager"
)

type TaxIncludedPricesJob struct {
	IOManager         filemanager.FileManager `json:"-"` // ignore this key
	TaxRate           float64                 `json:"lol_tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func NewTaxIncludedPricesJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{100, 200, 300},
	}
}

func (job *TaxIncludedPricesJob) LoadData() {

	lines, err := job.IOManager.ReactLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPricesJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		texIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", texIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}
