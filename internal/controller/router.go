package controller

import (
	"github.com/Orken1119/Register/pkg"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Orken1119/Register/internal/controller/auth_controller/auth"
	repository "github.com/Orken1119/Register/internal/repositories"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.DB

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/signin-as-volunteer", loginController.SigninAsVolunteer)
	router.POST("/signup-as-volunteer", loginController.SignupAsVolunteer)
	router.POST("/forgot-password", loginController.ForgotPassword)
	router.POST("/change-forgotten-password", loginController.ChangeForgottenPassword)
	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

}
