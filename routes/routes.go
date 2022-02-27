package routes

import (
	"FinalProjectGolangH8/auth"
	"FinalProjectGolangH8/comment"
	"FinalProjectGolangH8/handler"
	"FinalProjectGolangH8/photo"
	socialmedia "FinalProjectGolangH8/socialMedia"
	"FinalProjectGolangH8/user"
	"fmt"

	md "FinalProjectGolangH8/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	userRepository := user.NewRepository(db)
	photoRepository := photo.NewRepository(db)
	commentRepository := comment.NewRepository(db)
	socialMediaRepository := socialmedia.NewRepository(db)
	socialMediaRepository.FindAll()

	userService := user.NewService(userRepository)
	photoService := photo.NewService(photoRepository)
	commentService := comment.NewService(commentRepository)
	socialmediaService := socialmedia.NewService(socialMediaRepository, photoRepository)

	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	photoHandler := handler.NewPhotoHandler(photoService, userService)
	socialMediaHandler := handler.NewSocialMediaHandler(socialmediaService, userService)

	commentHandler := handler.NewCommentHandler(commentService, userService, photoService)

	r.GET("/foo", func(c *gin.Context) {
		fmt.Println("The URL: ", c.Request.URL.Path)
	})
	authz := md.AuthzMiddleware(photoService, commentService, socialmediaService)

	userRoute := r.Group("/users")
	userRoute.POST("/register", userHandler.RegisterUser)
	userRoute.POST("/login", userHandler.Login)
	userRoute.PUT("", md.AuthMiddleware(authService, userService), userHandler.UpdateUser)
	userRoute.DELETE("", md.AuthMiddleware(authService, userService), userHandler.DeleteUser)

	photoRoute := r.Group("/photos")
	photoRoute.POST("", md.AuthMiddleware(authService, userService), photoHandler.CreatePhoto)
	photoRoute.GET("", md.AuthMiddleware(authService, userService), photoHandler.GetAll)
	photoRoute.PUT("/:id", md.AuthMiddleware(authService, userService), authz, photoHandler.PutPhoto)
	photoRoute.DELETE("/:id", md.AuthMiddleware(authService, userService), authz, photoHandler.DeletePhoto)

	commentRoute := r.Group("/comments")
	commentRoute.POST("", md.AuthMiddleware(authService, userService), commentHandler.CreateComment)
	commentRoute.GET("", md.AuthMiddleware(authService, userService), commentHandler.GetAll)
	commentRoute.PUT("/:id", md.AuthMiddleware(authService, userService), authz, commentHandler.PutComment)
	commentRoute.DELETE("/:id", md.AuthMiddleware(authService, userService), authz, commentHandler.DeleteComment)

	socialMedia := r.Group("/socialmedias")
	socialMedia.POST("", md.AuthMiddleware(authService, userService), socialMediaHandler.CreateSocialMedia)
	socialMedia.GET("", md.AuthMiddleware(authService, userService), socialMediaHandler.GetAll)
	socialMedia.PUT("/:id", md.AuthMiddleware(authService, userService), authz, socialMediaHandler.PutSocialMedia)
	socialMedia.DELETE("/:id", md.AuthMiddleware(authService, userService), authz, socialMediaHandler.DeleteSocialMedia)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
