package main

import (
	"festApp/config"
	"festApp/controllers"
	"festApp/middleware"
	"festApp/repository"
	"festApp/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository   repository.UserRepository    = repository.NewUserRepository(db)
	artistRepository repository.ArtistRepository  = repository.NewArtistRepository(db)
	eventRepository  repository.EventRepository   = repository.NewEventRepository(db)
	jwtService       service.JWTService           = service.NewJWTService()
	userService      service.UserService          = service.NewUserService(userRepository)
	artistService    service.ArtistService        = service.NewArtistService(artistRepository)
	eventService     service.EventService         = service.NewEventService(eventRepository)
	authService      service.AuthService          = service.NewAuthService(userRepository)
	authController   controllers.AuthController   = controllers.NewAuthController(authService, jwtService)
	userController   controllers.UserController   = controllers.NewUserController(userService, jwtService)
	artistController controllers.ArtistController = controllers.NewArtistController(artistService, jwtService)
	eventController  controllers.EventController  = controllers.NewEventController(eventService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PATCH("/profile", userController.Update)
	}

	artistRoutes := r.Group("/artists", middleware.AuthorizeJWT(jwtService))
	{
		artistRoutes.POST("/", artistController.Insert)
		artistRoutes.PATCH("/:id", artistController.Update)
		artistRoutes.DELETE("/:id", artistController.Delete)
	}
	// Les routes get n'ont pas besoins de l'autorisation via token JWT
	getArtistRoutes := r.Group("/artists")
	{
		getArtistRoutes.GET("/", artistController.All)
		getArtistRoutes.GET("/:id", artistController.FindByID)
	}

	eventRoutes := r.Group("/events", middleware.AuthorizeJWT(jwtService))
	{
		eventRoutes.POST("/", eventController.Insert)
		eventRoutes.PATCH("/:id", eventController.Update)
		eventRoutes.DELETE("/:id", eventController.Delete)
	}
	// Les routes get n'ont pas besoins de l'autorisation via token JWT
	getEventRoutes := r.Group("/events")
	{
		getEventRoutes.GET("/", eventController.All)
		getEventRoutes.GET("/:id", eventController.FindByID)
	}

	r.Run()
}
