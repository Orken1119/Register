package main

import (
	"fmt"
	"log"

	controller "github.com/Orken1119/Register/internal/controller"
	pkg "github.com/Orken1119/Register/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDBConnection()

	ginRouter := gin.Default()

	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%d", 1140))
}
