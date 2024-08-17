package dto

type PostUrlDTO struct {
	Value string `json:"value" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}
