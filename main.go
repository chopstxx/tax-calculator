package main

import (
	"example.com/taxCalculator/cmdmanager"
	"example.com/taxCalculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0.09, 0.07, 0.1, 0.15}
	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates{
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob :=prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()

	}
	
	// fmt.Println(result)
}