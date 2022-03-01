package models

type Stories struct {
	Firstin  int    `json:"firstin"`
	Firstout int    `json:"firstout"`
	Flowin   int    `json:"flowin"`
	Flowout  int    `json:"flowout"`
	Rem      int    `json:"rem"`
	StoreId  int    `json:"store_id"`
	Name     string `json:"name"`
}