package models

type LinkMeta struct {
	BaseModel

	Entry       string  `json:"entry_id" gorm:"uniqueIndex"`
	Icon        string  `json:"icon"`
	URL         string  `json:"url"`
	Image       *string `json:"image"`
	Title       *string `json:"title"`
	Video       *string `json:"video"`
	Description *string `json:"description"`
}
