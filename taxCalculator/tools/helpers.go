package tools

import (
	"log"
	"github.com/drshott/tax-calculator/structures"
)

var taxSlabsNew = []structures.TaxSlab{
	{Range: "0-4L", Percentage: "0", Limit: 400000, Rate: 0, Tax: 0},
	{Range: "4L-8L", Percentage: "5%", Limit: 800000, Rate: 0.05, Tax: 20000},
	{Range: "8L-12L", Percentage: "10%", Limit: 1200000, Rate: 0.1, Tax: 60000},
	{Range: "12L-16L", Percentage: "15%", Limit: 1600000, Rate: 0.15, Tax: 120000},
	{Range: "16L-20L", Percentage: "20%", Limit: 2000000, Rate: 0.2, Tax: 200000},
	{Range: "20L-24L", Percentage: "25%", Limit: 2400000, Rate: 0.25, Tax: 300000},
	{Range: "24L+", Percentage: "30%", Limit: 1e9, Rate: 0.3, Tax: 1e9},
}

var taxSlabsOld = []structures.TaxSlab{
	{Range: "0-3L", Percentage: "0", Limit: 300000, Rate: 0, Tax: 0},
	{Range: "3L-7L", Percentage: "5%", Limit: 700000, Rate: 0.05, Tax: 20000},
	{Range: "7L-10L", Percentage: "10%", Limit: 1000000, Rate: 0.1, Tax: 50000},
	{Range: "10L-12L", Percentage: "15%", Limit: 1200000, Rate: 0.15, Tax: 80000},
	{Range: "12L-15L", Percentage: "20%", Limit: 1500000, Rate: 0.2, Tax: 140000},
	{Range: "15L+", Percentage: "30%", Limit: 1e9, Rate: 0.3, Tax: 1e9},
}

var taxSurcharge = []structures.TaxSlab{
	{Limit: 5000000, Rate: 0, Tax: 0},
	{Limit: 10000000, Rate: 0.1, Tax: 258000},
	{Limit: 20000000, Rate: 0.15, Tax: 558000},
	{Limit: 1e9, Rate: 0.25, Tax: 1e9},
}

func CalculateSlab(ammount float64, taxSlabs []structures.TaxSlab) (float64, []structures.RespTaxSlab) {
	var totalTax float64
	var prevLimit float64
	var prevTax float64
	var slabs []structures.RespTaxSlab
	
	for _, slab := range taxSlabs {
		if ammount > slab.Limit {
			slabs = append(slabs, structures.RespTaxSlab{
				Range: slab.Range,
				Rate: slab.Percentage,
				Amount: slab.Limit - prevLimit,
				SlabTax: slab.Tax,
			})
		} else {
			taxInSlab := (ammount - prevLimit) * slab.Rate
			totalTax = taxInSlab + prevTax
			slabs = append(slabs, structures.RespTaxSlab{
				Range: slab.Range,
				Rate: slab.Percentage,
				Amount: ammount - prevLimit,
				SlabTax: taxInSlab,
			})
			break
		}
		prevLimit = slab.Limit
		prevTax = slab.Tax
	}

	return totalTax, slabs

}

func CalculateTax(taxableIncome float64, newRegime bool) (float64, []structures.RespTaxSlab) {
	var taxSlabs []structures.TaxSlab
	if newRegime {
		taxSlabs = taxSlabsNew
	} else {
		taxSlabs = taxSlabsOld
	}
	totalTax, slab := CalculateSlab(taxableIncome, taxSlabs)
	return totalTax, slab
}

func CalculateCess(tax float64) float64 {
	cess := tax * 0.04
	return cess
}

func CalculateSurcharge(taxableIncome float64, tax float64, newRegime bool) float64 {
	if taxableIncome <= 5000000 {
		return 0
	}
	var surcharge float64
	var prevLimit float64
	for _, slab := range taxSurcharge {
		log.Printf("Slab limit: %f", slab.Limit)
		if taxableIncome < slab.Limit {
			surcharge = tax * slab.Rate
			log.Printf("Surcharge: %f", surcharge)
			log.Printf("Tax: %f", tax)
			limitTax, _ := CalculateTax(prevLimit, newRegime)
			log.Printf("Limit tax: %f", limitTax)
			log.Printf("tax income - prev limit: %f", taxableIncome - prevLimit)
			log.Printf("tax + surcharge - limit tax: %f", tax + surcharge - limitTax)
			if taxableIncome - prevLimit < (tax + surcharge) - limitTax {
				surcharge = (taxableIncome - prevLimit) - (tax - limitTax)
				log.Printf("Surcharge: %f", surcharge)
			}
			break
		}
		prevLimit = slab.Limit
	}
	return surcharge
}

func CalculateSurchargeRelief() {
//
}


func CalculateTaxRebate(taxableIncome float64, tax float64, newRegime bool) float64 {
	if newRegime {
		if taxableIncome <= 1200000 {
			return tax
		}
	} else {
		if taxableIncome <= 700000 {
			return tax
		}
	}
	return 0
}
