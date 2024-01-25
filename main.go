package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.String(http.StatusOK, "Hello, Go!")
}

func main() {
	router := gin.Default()

	router.GET("/albums/:id", getAlbum)
	router.GET("/albums", getAlbums)

	router.GET("/reduce/:id", priceReduction)

	router.GET("/", getHello)

	router.Run("localhost:8080")
}
