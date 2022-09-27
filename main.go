package main

import (
	"clean/config"
	"clean/config/infrastructure/datastore"
	"clean/config/infrastructure/registry"
	"clean/config/infrastructure/router"

	"github.com/gin-gonic/gin"
)

func main() {

	err := config.SetConfig()

	if err != nil {
		panic(err)
	}

	db, err := datastore.SetDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := registry.NewRegistry(db)

	g := gin.Default()
	router.SetRoute(r, g)
	g.Run(":9090")

}
