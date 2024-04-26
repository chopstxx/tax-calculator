package prices

import (
	"fmt"

	"example.com/taxCalculator/conversion"
	"example.com/taxCalculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"taxIncluded_prices"`
}
func (job *TaxIncludedPriceJob) LoadPrices(){

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	
	job.InputPrices = prices
	fmt.Println(prices)
	
}
func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}
	// fmt.Println(result)
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(fmt.Sprintf("result_%.0f.json", job.TaxRate*100),job)

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:     taxRate,
	}
}