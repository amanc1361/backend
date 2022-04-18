package modelsout

type InvoiceTax struct {
	SolarDate     string `json:"solar_date"`
	InvoiceNumner int    `json:"invoice_number"`
	PersonName    string `json:"person_name"`
	IncludeTax    int    `json:"include_tax"`
	NoIncludeTax  int    `json:"no_include_tax"`
	Tax           int    `json:"tax"`
	Description   string `json:"description"`
}