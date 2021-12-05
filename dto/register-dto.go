package dto

// RegisterDTO is used when client post a from /register url
type RegisterDTO struct {
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Firstname string `json:"firstname" form:"firstname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required"`
}
