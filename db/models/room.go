package models

type Room struct {
	Model
	Name string `json:"name"`
	Slug string `json:"slug"`
}
