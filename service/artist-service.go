package service

import (
	"festApp/dto"
	"festApp/entity"
	"festApp/repository"
	"github.com/mashingan/smapping"
	"log"
)

//ArtistService is a ....
type ArtistService interface {
	Insert(b dto.ArtistCreateDTO) entity.Artist
	Update(b dto.ArtistUpdateDTO) entity.Artist
	Delete(b entity.Artist)
	All() []entity.Artist
	FindByID(artistID uint64) entity.Artist
	//IsAllowedToEdit(userID string, artistID uint64) bool
}

type artistService struct {
	artistRepository repository.ArtistRepository
}

//NewArtistService .....
func NewArtistService(artistRepo repository.ArtistRepository) ArtistService {
	return &artistService{
		artistRepository: artistRepo,
	}
}

func (service *artistService) Insert(b dto.ArtistCreateDTO) entity.Artist {
	artist := entity.Artist{}
	err := smapping.FillStruct(&artist, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.artistRepository.InsertArtist(artist)
	return res
}

func (service *artistService) Update(b dto.ArtistUpdateDTO) entity.Artist {
	artist := entity.Artist{}

	err := smapping.FillStruct(&artist, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.artistRepository.UpdateArtist(artist)
	return res
}

func (service *artistService) Delete(b entity.Artist) {
	service.artistRepository.DeleteArtist(b)
}

func (service *artistService) All() []entity.Artist {
	return service.artistRepository.AllArtist()
}

func (service *artistService) FindByID(artistID uint64) entity.Artist {
	return service.artistRepository.FindArtistByID(artistID)
}

//func (service *artistService) IsAllowedToEdit(userID string, artistID uint64) bool {
//	//b := service.artistRepository.FindArtistByID(artistID)
//	//id := fmt.Sprintf("%v", b.UserID)
//	return service.artistRepository.FindArtistByID(artistID)
//}
