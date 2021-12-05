package repository

import (
	"festApp/entity"
	"gorm.io/gorm"
)

//ArtistRepository is a ....
type ArtistRepository interface {
	InsertArtist(b entity.Artist) entity.Artist
	UpdateArtist(b entity.Artist) entity.Artist
	DeleteArtist(b entity.Artist)
	AllArtist() []entity.Artist
	FindArtistByID(artistID uint64) entity.Artist
}

type artistConnection struct {
	connection *gorm.DB
}

//NewArtistRepository creates an instance ArtistRepository
func NewArtistRepository(dbConn *gorm.DB) ArtistRepository {
	return &artistConnection{
		connection: dbConn,
	}
}

func (db *artistConnection) InsertArtist(b entity.Artist) entity.Artist {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *artistConnection) UpdateArtist(b entity.Artist) entity.Artist {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *artistConnection) DeleteArtist(b entity.Artist) {
	db.connection.Delete(&b)
}

func (db *artistConnection) FindArtistByID(artistID uint64) entity.Artist {
	var artist entity.Artist
	db.connection.Preload("User").Find(&artist, artistID)
	return artist
}

func (db *artistConnection) AllArtist() []entity.Artist {
	var artists []entity.Artist
	db.connection.Preload("User").Find(&artists)
	return artists
}
