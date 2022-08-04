package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

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
		status := 200
		msg := ""
		data, err := get(url)
		if err == nil {
			msg = data
		} else {
			status = 500
			msg = err.Error()
		}
		c.JSON(status, gin.H{
			"message": msg,
		})
	}
	return gin.HandlerFunc(fn)
}
func get(url string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("lbonly", "yes")
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
