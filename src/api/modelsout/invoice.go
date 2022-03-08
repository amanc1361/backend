package modelsout

type Invoice struct {
	Id               int    `json:"id"`
	CustomerName     string `json:"customer_name"`
	InvoiceNumber    int    `json:"invoice_number"`
	InvoiceBuynumber string `json:"buy_number"`
	SolarDate        string `json:"solar_date"`
	Amount           int    `json:"amount"`
	Description      string `json:"description"`
}