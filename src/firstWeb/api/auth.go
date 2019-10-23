package api

import (
	"firstWeb/models"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Engine) {
	c.GET("/auth", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")

		isOK := models.CheckAuth(username, password)
		if !isOK {
			c.JSON(http.StatusOK, gin.H{"message": "authError"})
		}

		token, err := util.GenerateToken(username, password)

		if err != nil {
			//...
		}
		c.JSON(http.StatusOK, gin.H{"tokon": token})
	})

}
