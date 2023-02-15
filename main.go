package main

import (
	"api-data-mhs/auth"
	"api-data-mhs/db"
	"api-data-mhs/handler"
	"api-data-mhs/mhs"
	"api-data-mhs/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDB()

	userRepository := user.NewRepository(db)
	mhsRepository := mhs.NewRepository(db)

	userService := user.NewService(userRepository)
	mhsService := mhs.NewService(mhsRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	mhsHandler := handler.NewMhsHandler(mhsService)

	router := gin.Default()
	api := router.Group("/api-mhs")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.PUT("/forgot-password/:nim", userHandler.ForgotPassword)

	api.POST("/create", auth.Middleware(authService, userService), mhsHandler.Create)
	api.PUT("/update/:nim", auth.Middleware(authService, userService), mhsHandler.Update)
	api.DELETE("/delete/:nim", auth.Middleware(authService, userService), mhsHandler.Delete)
	api.GET("/mahasiswa", auth.Middleware(authService, userService), mhsHandler.FindAll)
	api.GET("/mahasiswa/:nim", auth.Middleware(authService, userService), mhsHandler.FindByNim)

	router.Run()

}
