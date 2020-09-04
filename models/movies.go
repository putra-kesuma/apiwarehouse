package models

type Movies struct {
	IdMovie int `json:"id"`
	Title string `json:"title"`
	Duration string `json:"duration"`
	ImageUrl string `json:"imageUrl"`
	Synopsis string `json:"synopsis"`
}