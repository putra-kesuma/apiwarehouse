package models

type Item struct {
	IdItem int `json:"id"`
	Name string `json:"nameitem"`
	Dimension int `json:"dimension"`
	IdTypeItem int `json:"idtypeitem"`
	NameItem string `json:"nameItem"`
}
