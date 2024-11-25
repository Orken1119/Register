package controller

import (
	"time"

	"github.com/Orken1119/Register/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Orken1119/Register/internal/controller/auth_controller/auth"
	"github.com/Orken1119/Register/internal/controller/auth_controller/ivent"
	"github.com/Orken1119/Register/internal/controller/auth_controller/middleware"
	"github.com/Orken1119/Register/internal/controller/auth_controller/user"
	repository "github.com/Orken1119/Register/internal/repositories"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.DB

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}

	iventController := &ivent.IventController{
		IventRepository: repository.NewIventRepository(db),
	}

	// CORS settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Укажите URL вашего фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authenticationRouter := router.Group("/authentication")
	{
		authenticationRouter.POST("/signup-as-volunteer", loginController.SignupAsVolunteer)
		authenticationRouter.POST("/signin-as-volunteer", loginController.SigninAsVolunteer)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/forgot-password", loginController.ForgotPassword)
	router.POST("/change-forgotten-password", loginController.ChangeForgottenPassword)
	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

	router.Use(middleware.JWTAuth(`access-secret-key`))

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
		userRouter.PUT("/edit-profile", userController.EditPersonalData)
		userRouter.PUT("/change-password", userController.ChangePassword)
	}

	iventRouter := router.Group("/ivents")
	{
		iventRouter.POST("/create-ivent", iventController.CreateIvent)
		iventRouter.GET("/get-ivents", iventController.GetAllIvent)
		iventRouter.GET("/get-ivent-by-id", iventController.GetIventById)
		iventRouter.PUT("/update-ivent", iventController.UpdateIvent)
		iventRouter.DELETE("/delete-ivent", iventController.DeleteIvent)
	}

}
