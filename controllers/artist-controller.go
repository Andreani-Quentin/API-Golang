package controllers

import (
	"festApp/dto"
	"festApp/entity"
	"festApp/helper"
	"festApp/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//ArtistController is a ...
type ArtistController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type artistController struct {
	artistService service.ArtistService
	jwtService  service.JWTService
}

//NewArtistController create a new instances of ArtistsController
func NewArtistController(artistServ service.ArtistService, jwtServ service.JWTService) ArtistController {
	return &artistController{
		artistService: artistServ,
		jwtService:  jwtServ,
	}
}

func (c *artistController) All(context *gin.Context) {
	var artists []entity.Artist = c.artistService.All()
	res := helper.BuildResponse(true, "OK", artists)
	context.JSON(http.StatusOK, res)
}

func (c *artistController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var artist entity.Artist = c.artistService.FindByID(id)
	if (artist == entity.Artist{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", artist)
		context.JSON(http.StatusOK, res)
	}
}

func (c *artistController) Insert(context *gin.Context) {
	var artistCreateDTO dto.ArtistCreateDTO
	errDTO := context.ShouldBind(&artistCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		//err := context.SaveUploadedFile(artistCreateDTO.Image,"assets/"+ artistCreateDTO.Image.Filename)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		// À checker
		fmt.Println(convertedUserID)
		if err == nil {
			result := c.artistService.Insert(artistCreateDTO)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}
}

func (c *artistController) Update(context *gin.Context) {
	var artistUpdateDTO dto.ArtistUpdateDTO
	errDTO := context.ShouldBind(&artistUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if userID != "" {
		result := c.artistService.Update(artistUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *artistController) Delete(context *gin.Context) {
	var artist entity.Artist
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	artist.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if userID != "" {
		c.artistService.Delete(artist)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *artistController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}




///*	GET /artists
///*	Get all artists
///* ********************************************** */
//func FindArtists(c *gin.Context) {
//	var artists []entity.Artist
//	config.DB.Find(&artists)
//
//	c.JSON(http.StatusOK, gin.H{"data": artists})
//}
//
//
///*	GET /artists/:id
///*	Get one artist
///* ********************************************** */
//func FindArtist(c *gin.Context) { // trouve un model si il existe
//	var artist entity.Artist
//
//	if err := config.DB.Where("id = ?", c.Param("id")).First(&artist).Error; err !=nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error" : "Livre non trouvé"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"data": artist})
//}
//
//
///*	POST /artists/
///*	Create new artist
///* ********************************************** */
//func CreateArtist(c *gin.Context) {
//	// Validate input
//	var input dto.CreateArtistInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Create artist
//	artist := entity.Artist{Name: input.Name, Description: input.Description}
//	config.DB.Create(&artist)
//
//	c.JSON(http.StatusOK, gin.H{"data": artist})
//}
//
//
///*	PATCH /artists/:id
///*	Update a artist
///* ********************************************** */
//func UpdateArtist(c *gin.Context) {
//	// Get model if exist
//	var artist entity.Artist
//	if err := config.DB.Where("id = ?", c.Param("id")).First(&artist).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
//		return
//	}
//
//	// Validate input
//	var input dto.UpdateArtistInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	config.DB.Model(&artist).Updates(input)
//
//	c.JSON(http.StatusOK, gin.H{"data": artist})
//}
//
///*	DELETE /artists/:id
///*	Delete a artist
///* ********************************************** */
//func DeleteArtist(c *gin.Context) {
//	//Get model if existed
//	var artist entity.Artist
//	if err := config.DB.Where("id = ?", c.Param("id")).First(&artist).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !"})
//		return
//	}
//
//	config.DB.Delete(&artist)
//
//	c.JSON(http.StatusOK, gin.H{"data": true})
//}
//
////func Var_dump(expression ...interface{} ) {
////	fmt.Println(fmt.Sprintf("%#v", expression))
////}