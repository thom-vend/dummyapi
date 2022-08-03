package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	lburl := os.Getenv("LBURL")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.GET("/hitme", hitmeHandler(lburl))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}

func hitmeHandler(url string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		status := 500
		msg := "error"
		data, err := http.Get(url)
		if err == nil {
			body, err := ioutil.ReadAll(data.Body)
			if err == nil {
				status = 200
				msg = string(body)
			}
		}
		c.JSON(status, gin.H{
			"message": msg,
		})
	}
	return gin.HandlerFunc(fn)
}
