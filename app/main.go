package app

import "github.com/gin-gonic/gin"

func GetData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hello go!",
	})
}

func main() {
	server := gin.Default()
	server.GET("/", GetData)
}
