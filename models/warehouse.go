package models

type Warehouse struct{
	IdWarehouse int `json:"id"`
	Name string `json:"namewarehouse"`
	Capacity int `json:"capacity"`
	Address string `json:"address"`
	IdTypeWarehouse int `json:"idtypewarehouse"`
	NameType string `json:"nametype"`
}

