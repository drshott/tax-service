package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/drshott/tax-calculator/structures"
	"github.com/drshott/tax-calculator/tools"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/static"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/index.html")
	}).Methods("GET")

	router.HandleFunc("/api/tax/", TaxHandler).Methods("GET").Queries("income", "{income}")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func TaxHandler(w http.ResponseWriter, r *http.Request) {
	// var requestData map[string]interface{}

	// Decode JSON body into a map
	// err := json.NewDecoder(r.Body).Decode(&requestData)
	// if err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	//	return
	// }

	queryParams := r.URL.Query()
	income := queryParams.Get("income")

	totalIncome, err := strconv.ParseFloat(income, 64)
	if err != nil {
		// If the conversion fails, return an error
		http.Error(w, "Invalid income value", http.StatusBadRequest)
		return
	}
	log.Printf("Income received: %f", totalIncome)
	// totalIncome := requestData["income"].(float64)
	taxableIncome := totalIncome - 75000.000000
	log.Printf("Taxable income: %f", taxableIncome)

	totalTaxNew, slabsNew := tools.CalculateTax(taxableIncome, true)
	totalTaxOld, slabsOld := tools.CalculateTax(taxableIncome, false)
	log.Printf("Total tax 25: %f", totalTaxNew)
	log.Printf("Total tax 24: %f", totalTaxOld)

	surchargeNew := tools.CalculateSurcharge(taxableIncome, totalTaxNew, true)
	surchargeOld := tools.CalculateSurcharge(taxableIncome, totalTaxOld, false)

	rebateNew := tools.CalculateTaxRebate(taxableIncome, totalTaxNew, true)
	rebateOld := tools.CalculateTaxRebate(taxableIncome, totalTaxOld, false)

	taxNew := totalTaxNew - rebateNew + surchargeNew
	taxOld := totalTaxOld - rebateOld + surchargeOld

	cessNew := tools.CalculateCess(taxNew)
	cessOld := tools.CalculateCess(taxOld)

	actualTaxNew := taxNew + cessNew + 2400
	actualTaxOld := taxOld + cessOld + 2400
	
	log.Printf("Actual tax 25: %f", actualTaxNew)
	log.Printf("Actual tax 24: %f", actualTaxOld)

	//jsonResp := respJson{
	//	Tax25: fmt.Sprintf("%.2f", actualTax25),
	//	Tax24: fmt.Sprintf("%.2f", actualTax24),
	//}
	jsonResp := structures.RespJson{
		TaxNew: structures.RespTax{
			Tax: totalTaxNew,
			Surcharge: surchargeNew,
			Cess: cessNew,
			Rebate: rebateNew,
			TotalTax: actualTaxNew,
			Slabs: slabsNew,
		},
		TaxOld: structures.RespTax{
			Tax: totalTaxOld,
			Surcharge: surchargeOld,
			Cess: cessOld,
			Rebate: rebateOld,
			TotalTax: actualTaxOld,
			Slabs: slabsOld,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonResp)
}

