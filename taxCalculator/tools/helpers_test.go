package tools

import (
	"testing"
)

func TestCalculateSlab(t *testing.T) {
	tests := []struct {
		income    float64
		newRegime bool
		expected  float64
	}{
		{500000, true, 5000},  // Under new regime
		{500000, false, 10000}, // Under old regime
		{1500000, true, 105000},
		{1500000, false, 140000},
	}

	for _, tt := range tests {
		tax, _ := CalculateTax(tt.income, tt.newRegime)
		if tax != tt.expected {
			t.Errorf("CalculateTax(%f, %t) = %f; expected %f", tt.income, tt.newRegime, tax, tt.expected)
		}
	}
}

func TestCalculateCess(t *testing.T) {
	tax := 100000
	expected := 4000.0
	cess := CalculateCess(float64(tax))
	if cess != expected {
		t.Errorf("CalculateCess(%f) = %f; expected %f", float64(tax), cess, expected)
	}
}

func TestCalculateSurcharge(t *testing.T) {
	tests := []struct {
		income    float64
		tax       float64
		newRegime bool
		expected  float64
	}{
		{6000000, 1510000, true, 151000},
		{11000000, 3010000, false, 451500},
		{21000000, 6010000, true, 570000},
	}

	for _, tt := range tests {
		surcharge := CalculateSurcharge(tt.income, tt.tax, tt.newRegime)
		if surcharge != tt.expected {
			t.Errorf("CalculateSurcharge(%f, %f, %t) = %f; expected %f", tt.income, tt.tax, tt.newRegime, surcharge, tt.expected)
		}
	}
}

func TestCalculateTaxRebate(t *testing.T) {
	tests := []struct {
		income    float64
		tax       float64
		newRegime bool
		expected  float64
	}{
		{600000, 10000, false, 10000},
		{1300000, 50000, true, 0},
		{650000, 12000, false, 12000},
	}

	for _, tt := range tests {
		rebate := CalculateTaxRebate(tt.income, tt.tax, tt.newRegime)
		if rebate != tt.expected {
			t.Errorf("CalculateTaxRebate(%f, %f, %t) = %f; expected %f", tt.income, tt.tax, tt.newRegime, rebate, tt.expected)
		}
	}
}
