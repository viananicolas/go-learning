package main

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "www.pitchshifter.com", Artist: "Pitchshifter", Price: 56.99},
	{ID: "2", Title: "Selfless", Artist: "Godflesh", Price: 17.99},
	{ID: "3", Title: "Rio Grande Blood", Artist: "Ministry", Price: 39.99},
}

func main() {
	fmt.Println("Howdy, partner!")

	request := gin.Default()

	request.GET("/albums", getAlbums)
	request.GET("/albums/:id", getAlbum)
	request.POST("/albums", postAlbum)
	request.DELETE("/albums/:id", deleteAlbum)

	request.Run()
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	index := sort.Search(len(albums), func(i int) bool {
		return string(albums[i].ID) >= id
	})

	if index < len(albums) && albums[index].ID == id {
		selectedAlbum := albums[index]
		c.IndentedJSON(http.StatusOK, selectedAlbum)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	sort.Slice(albums, func(i, j int) bool { return albums[i].ID < albums[j].ID })
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	index := sort.Search(len(albums), func(i int) bool {
		return string(albums[i].ID) >= id
	})

	if index < len(albums) && albums[index].ID == id {
		albums = append(albums[:index], albums[index+1:]...)
		c.Status(http.StatusNoContent)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
