package app

import (
	"PauloLeal/go-gin-app-template/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var appInstance *application
var appOnce sync.Once

type application struct {
	router *gin.Engine
}

func App() *application {
	appOnce.Do(func() {
		g := gin.New()
		g.Use(gin.Recovery())

		appInstance = &application{router: g}
		appInstance.createRoutes()
	})
	return appInstance
}

func (app *application) RunServer(port int) error {
	return app.router.Run(fmt.Sprintf(":%d", port))
}

func (app *application) createRoutes() {
	healthController := controllers.NewHealthController()
	app.router.GET("/health", healthController.Get)
}
