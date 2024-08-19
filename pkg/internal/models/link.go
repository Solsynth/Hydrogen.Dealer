package models

type LinkMeta struct {
	BaseModel

	Icon        string  `json:"icon"`
	URL         string  `json:"url"`
	Image       *string `json:"image"`
	Title       *string `json:"title"`
	Video       *string `json:"video"`
	Description *string `json:"description"`
}
