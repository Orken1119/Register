package controller

import (
	"github.com/Orken1119/Register/internal/controller/auth_controller/middleware"
	"github.com/Orken1119/Register/pkg"
	"github.com/gin-gonic/gin"

	"github.com/Orken1119/Register/internal/controller/auth_controller/auth"
	repository "github.com/Orken1119/Register/internal/repositories"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.DB

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

	router.Use(middleware.JWTAuth(`access-secret-key`))

}
