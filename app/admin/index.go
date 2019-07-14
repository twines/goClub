package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "/admin/index/index.html", gin.H{})
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "/admin/index/login.html", gin.H{})
}
func DoLogin(c *gin.Context) {
	//c.Request.URL.Path= "/admin/index"
	c.Redirect(http.StatusMovedPermanently, "/admin/index")
}
