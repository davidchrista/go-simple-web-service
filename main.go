package main

import (
	"net/http"
	"time"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
)

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (album *Album) reducePrice(percent int) {
	album.Price *= 1.0 - float64(percent)/100.0
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func findById(albums []Album, id string) Album {
	for _, a := range albums {
		if a.ID == id {
			return a
		}
	}
	return Album{}
}

func findByIdMut(albums *[]Album, id string) *Album {
	a := *albums
	for i := range a {
		if a[i].ID == id {
			return &a[i]
		}
	}
	return nil
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	a := findById(albums, id)
	c.IndentedJSON(http.StatusOK, a)
}

func priceReduction(c *gin.Context) {
	id := c.Param("id")
	a := findByIdMut(&albums, id)
	if a != nil {
		a.reducePrice(10)
	}
}

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Go!")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := "https://dev-gzm0pgbh.us.auth0.com/"

		config := auth0.NewConfiguration(
			auth0.NewJWKClient(auth0.JWKClientOptions{URI: domain + ".well-known/jwks.json"}, nil),
			[]string{"http://localhost:4000"},
			domain,
			jose.RS256,
		)

		validator := auth0.NewValidator(config, nil)

		token, err := validator.ValidateRequest(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", token)

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(AuthMiddleware())

	router.GET("/albums/:id", getAlbum)
	router.GET("/albums", getAlbums)

	router.GET("/reduce/:id", priceReduction)

	router.GET("/", getHello)

	router.Run("localhost:4200")
}
