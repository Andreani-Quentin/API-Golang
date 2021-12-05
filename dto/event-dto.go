package dto

type EventCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Hour        string `json:"hour" form:"hour" binding:"required"`
}

type EventUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Hour        string `json:"hour" form:"hour" binding:"required"`
}
