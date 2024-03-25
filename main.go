package main

import (
	"mygram-finalprojectdts/controller"
	"mygram-finalprojectdts/lib"
	"mygram-finalprojectdts/middleware"
	"mygram-finalprojectdts/model"
	"mygram-finalprojectdts/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.InitDatabase()

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	photoRepository := repository.NewPhotoRepository(db)
	photoController := controller.NewPhotoController(photoRepository)

	commentRepository := repository.NewCommentRepository(db)
	commentController := controller.NewCommentController(commentRepository)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaController := controller.NewSocialMediaController(socialMediaRepository)

	ginEngine := gin.Default()

	//routes user

	ginEngine.POST("users/register", userController.Register)
	ginEngine.POST("users/login", userController.Login)

	userGroup := ginEngine.Group("/users", middleware.AuthMiddleware)

	userGroup.PUT("/", userController.Update)
	userGroup.DELETE("/", userController.Delete)

	//routes photos

	photoGroup := ginEngine.Group("/photos", middleware.AuthMiddleware)

	photoGroup.POST("/", photoController.Create)
	photoGroup.GET("/", photoController.GetAll)
	photoGroup.GET("/:photoId", photoController.GetOne)
	photoGroup.PUT("/:photoId", middleware.PhotoAuth(db), photoController.Update)
	photoGroup.DELETE("/:photoId", middleware.PhotoAuth(db), photoController.Delete)

	//routes comment

	commentGroup := ginEngine.Group("/comments", middleware.AuthMiddleware)

	commentGroup.POST("/", commentController.Create)
	commentGroup.GET("/", commentController.GetAll)
	commentGroup.GET("/:commentId", commentController.GetOne)
	commentGroup.PUT("/:commentId", middleware.CommentAuth(db), commentController.Update)
	commentGroup.DELETE("/:commentId", middleware.CommentAuth(db), commentController.Delete)

	//routes socmed

	socialMediaGroup := ginEngine.Group("/socialmedias", middleware.AuthMiddleware)

	socialMediaGroup.POST("/", socialMediaController.Create)
	socialMediaGroup.GET("/", socialMediaController.GetAll)
	socialMediaGroup.GET("/:socialmediaId", socialMediaController.GetOne)
	socialMediaGroup.PUT("/:socialmediaId", middleware.SocialMediaAuth(db), socialMediaController.Update)
	socialMediaGroup.DELETE("/:socialmediaId", middleware.SocialMediaAuth(db), socialMediaController.Delete)

	err = ginEngine.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
