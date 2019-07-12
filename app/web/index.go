package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	c.HTML(http.StatusOK, "/index/index.html", gin.H{
		"title": "Users",
	})
}
func Player(c *gin.Context) {
	c.HTML(http.StatusOK, "/player/index.html", gin.H{
		"title": "Users",
	})
}
