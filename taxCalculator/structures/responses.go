package structures

type RespJson struct {
	TaxNew RespTax `json:"taxNew"`
	TaxOld RespTax `json:"taxOld"`
}

type RespTax struct {
	Tax float64 `json:"tax"`
	Surcharge float64 `json:"surcharge"`
	Cess float64 `json:"cess"`
	Rebate float64 `json:"rebate"`
	TotalTax float64 `json:"totalTax"`
	Slabs []RespTaxSlab `json:"slabs"`
}

type RespTaxSlab struct {
	Range string `json:"range"`
	Rate string `json:"rate"`
	Amount float64 `json:"amount"`
	SlabTax float64 `json:"slabTax"`
}