package app 

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() (*App, error){
	router := gin.Default()

	app := &App{
		router: router,
	}

	app.setupRoutes()

	return app,nil
}

func (app *App) setupRoutes(){
	v1 := app.router.Group("/api/v1")
	{
		v1.GET("/health", app.handleHealth)
	}
}

func (app *App) Run() error {
	return app.router.Run(":8080")
}

func (app *App) handleHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}