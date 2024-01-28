package main

import (
	"github.com/davidchrista/go-simple-web-service/handlers"
	"github.com/davidchrista/go-simple-web-service/mw"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(mw.CorsMiddleware()).Use(mw.AuthMiddleware())

	router.GET("/albums/:id", handlers.GetAlbum)
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/reduce/:id", handlers.PriceReduction)

	router.GET("/", handlers.GetHello)

	router.Run("localhost:4200")
}
