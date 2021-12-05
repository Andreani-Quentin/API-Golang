package entity

type Event struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Hour        string `json:"hour"`
}
