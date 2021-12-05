package dto

// UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Firstname string `json:"firstname" form:"firstname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password,omitempty" form:"password,omitempty"`
}
