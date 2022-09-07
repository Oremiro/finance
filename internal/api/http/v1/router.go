package v1

import (
	"github.com/gin-gonic/gin"
)

type (
	IController interface {
		AddRoutes(routerGroup *gin.RouterGroup)
	}
)

func UseRouter(handler *gin.Engine, services []IController) {

	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	routerGroup := handler.Group("/v1")

	for _, service := range services {
		service.AddRoutes(routerGroup)
	}

}
