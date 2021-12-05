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

//EventController is a ...
type EventController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type eventController struct {
	eventService service.EventService
	jwtService  service.JWTService
}

//NewEventController create a new instances of EventsController
func NewEventController(eventServ service.EventService, jwtServ service.JWTService) EventController {
	return &eventController{
		eventService: eventServ,
		jwtService:  jwtServ,
	}
}

func (c *eventController) All(context *gin.Context) {
	var events []entity.Event = c.eventService.All()
	res := helper.BuildResponse(true, "OK", events)
	context.JSON(http.StatusOK, res)
}

func (c *eventController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var event entity.Event = c.eventService.FindByID(id)
	if (event == entity.Event{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", event)
		context.JSON(http.StatusOK, res)
	}
}

func (c *eventController) Insert(context *gin.Context) {
	var eventCreateDTO dto.EventCreateDTO
	errDTO := context.ShouldBind(&eventCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		// À checker
		fmt.Println(convertedUserID)
		if err == nil {
			result := c.eventService.Insert(eventCreateDTO)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}
}

func (c *eventController) Update(context *gin.Context) {
	var eventUpdateDTO dto.EventUpdateDTO
	errDTO := context.ShouldBind(&eventUpdateDTO)
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
		result := c.eventService.Update(eventUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *eventController) Delete(context *gin.Context) {
	var event entity.Event
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	event.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if userID != "" {
		c.eventService.Delete(event)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *eventController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}