package models

type Pageinate struct {
	FirstPage    string `json:"first_page"`
	LastPage     string `json:"last_page"`
	NextPage     string `json:"next_page"`
	PreviousPage string `json:"pre_page"`
}
