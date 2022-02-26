package modelsout

type Kradex struct {
	StoreActionID   int64   `json:"store_action_id"`
	StoreObjectID   int64   `json:"store_object_id"`
	StoreActionCode int64   `json:"store_action_code"`
	Countin         float32 `json:"countin"`
	Price           int     `json:"price"`
	Countout        float32 `json:"countout"`
	Description     string  `json:"description"`
	StoreObjectCode int64   `json:"store_object_code"`
	StoreObjecName  string  `json:"store_object_name"`
	SolarDate       string  `json:"solar_date"`
	UnitName        string  `json:"unit_name"`
}
