package router

import (
	"clean/config/infrastructure/registry"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *registry.Registry, g *gin.Engine) {

	g.GET("/users", r.NewUserRegistry().FindAll)
	g.POST("/users", r.NewUserRegistry().Create)
	g.GET("/user/:id", r.NewUserRegistry().Find)
	g.DELETE("/users/:id", r.NewUserRegistry().Delete)

}
