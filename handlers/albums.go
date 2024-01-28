package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (album *Album) reducePrice(percent int) {
	album.Price *= 1.0 - float64(percent)/100.0
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbum(c *gin.Context) {
	id := c.Param("id")
	a := findById(albums, id)
	c.IndentedJSON(http.StatusOK, a)
}

func PriceReduction(c *gin.Context) {
	id := c.Param("id")
	a := findByIdMut(&albums, id)
	if a != nil {
		a.reducePrice(10)
	}
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
