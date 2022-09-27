package router

import (
	"clean/config/infrastructure/registry"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *registry.Registry, g *gin.Engine) {

	api := g.Group("/api")
	{
		api.GET("/users", r.NewUserRegistry().FindAll)
		api.POST("/users", r.NewUserRegistry().Create)
		api.GET("/user/:id", r.NewUserRegistry().Find)
		api.DELETE("/users/:id", r.NewUserRegistry().Delete)
	}

}
