package structures

type TaxSlab struct {
	Limit float64
	Rate  float64
	Tax  float64
	Range string
	Percentage string
}

type Regime struct {
	Range string
	Rate  string
	Max   string
}