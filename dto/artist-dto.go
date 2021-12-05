package dto

type ArtistCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Image       string `json:"image" form:"image"`
}

type ArtistUpdateDTO struct {
	ID          uint64   `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Image       string `json:"image" form:"image"`
}
