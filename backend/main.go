package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		nums := []int{1, 2, 3, 4}
		c.JSON(http.StatusOK, nums)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"action": "pong",
		})
	})
	r.GET("/pong", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"action": "ping",
		})
	})

	log.Fatal(r.Run())
}
