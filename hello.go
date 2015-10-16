package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/stringutil"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", stringutil.Reverse(name))
	})

	router.Run(":8080")
}
