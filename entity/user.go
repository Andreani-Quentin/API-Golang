package entity

type User struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	LastName  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  string `json:"password" gorm:"->;<-;not null"`
	Token     string `json:"token,omitempty" gorm:"-"`
}
