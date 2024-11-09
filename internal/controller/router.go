package controller

import (
	"time"

	"github.com/Orken1119/Register/pkg"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Укажите URL вашего фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/signin-as-volunteer", loginController.SigninAsVolunteer)
	router.POST("/signup-as-volunteer", loginController.SignupAsVolunteer)
	router.POST("/forgot-password", loginController.ForgotPassword)
	router.POST("/change-forgotten-password", loginController.ChangeForgottenPassword)
	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

}
