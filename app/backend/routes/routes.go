package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/songs", controllers.GetSongs)
	router.POST("/api/songs", controllers.AddSong)
}
