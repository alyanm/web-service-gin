package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID	 string `json:"id"`
	Title	 string `json:"title"`
	Artist	 string `json:"artist"`
	Price	 float32 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumByID)
	router.Run("localhost:8080")
}

/* getAlbums responds with the list of all albums as JSON.
test:
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
	*/

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

/* postAlbums adds an album from JSON received in the request body.
test:
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
	*/

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/* getAlbumsByID responds with the album with the matching ID.
test:
curl http://localhost:8080/albums/2
*/
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter. */
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

/* updateAlbumByID updates an album from JSON received in the request body.
test:
curl http://localhost:8080/albums/2 \
	--include \
	--header "Content-Type: application/json" \
	--request "PUT" \
	--data '{"id": "2","title": "Jeru","artist": "Gerry Mulligan","price": 19.99}'
	*/
func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON provided"});
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}