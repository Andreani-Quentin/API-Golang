package entity

type Artist struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
