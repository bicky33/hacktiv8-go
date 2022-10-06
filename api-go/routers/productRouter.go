package routers

import (
	"api-go/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/albums", controllers.SeedAlbum)
	router.GET("/album/:title", controllers.ShowSingle)
	return router
}
