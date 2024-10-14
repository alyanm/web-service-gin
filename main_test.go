package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alyanm/web-service-gin/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db.InitDB("root:chicchan@tcp(127.0.0.1:3306)/albumdb")
	db.InitializeTestData()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	return router
}

func TestGetAlbums(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Sarah Vaughan")
	log.Println(w.Body.String())
}

func TestPostAlbums(t *testing.T) {
	router := setupRouter()

	/** Delete previous test run data **/
	req, _ := http.NewRequest("DELETE", "/albums/4", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	newAlbum := `{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}`
	w2 := httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/albums", bytes.NewBufferString(newAlbum))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w2, req)

	assert.Equal(t, http.StatusCreated, w2.Code)
	assert.Contains(t, w2.Body.String(), "The Modern Sound of Betty Carter")
}

func TestGetAlbumsByID(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Blue Train")
}

func TestUpdateAlbumByID(t *testing.T) {
	router := setupRouter()

	updatedAlbum := `{"id": "2", "title": "Jeru", "artist": "Gerry Mulligan", "price": 19.99}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/albums/2", bytes.NewBufferString(updatedAlbum))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Jeru")
}

func TestDeleteAlbumByID(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "album deleted")
}
