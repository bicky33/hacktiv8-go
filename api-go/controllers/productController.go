package controllers

import (
	"api-go/datastruct"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var StoredAlbums []datastruct.Album

func SeedAlbum(ctx *gin.Context) {
	var newAlbum datastruct.Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	StoredAlbums = append(StoredAlbums, newAlbum)

	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func ShowALl(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, StoredAlbums)
}

func ShowSingle(ctx *gin.Context) {
	dataTitle := ctx.Param("title")
	var newAlbum datastruct.Album
	var isExist bool

	for _, album := range StoredAlbums {
		if album.Tittle == dataTitle {
			isExist = true
			newAlbum = album
			break
		}
	}

	if !isExist {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprint("not found album with title %w", dataTitle),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, newAlbum)

}
