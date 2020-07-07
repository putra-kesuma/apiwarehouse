package models


type WarehouseStorage struct{
	IdWarehouseStorage int `json:"id"`
	IdWarehouse int `json:"idwarehouse"`
	NameWarehouse string `json:"namewarehouse"`
	Capacity int `json:"capacity"`
	Address string `json:"address"`
	NameTypeWarehouse string `json:"nametypewarehouse"`
	IdItem int `json:"iditem"`
	NameItem string `json:"nameitem"`
	Dimension int `json:"dimension"`
	NameTypeItem string `json:"nametypeitem"`
}
