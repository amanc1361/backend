package models

type StoreRemObjects struct {
	Firstin        float32 `json:"firstin"`
	Flowin         float32 `json:"flowin"`
	Flowout        float32 `json:"flowout"`
	Lrem           float32 `json:"lrem"`
	FstoreId       int     `json:"fstore_id"`
	Fname          string  `json:"fname"`
	LstoreId       int     `json:"lstore_id"`
	Lname          string  `json:"lname"`
	FstoreObjectId int     `json:"fstore_object_id"`
	Fcode          int     `json:"fcode"`
	LstoreObjectId int     `json:"lstore_object_id"`
	Lcode          int     `json:"lcode"`
}