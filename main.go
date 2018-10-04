package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

type PostData struct {
	Hello string `json:"hello"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		var json PostData
		c.BindJSON(&json)
		id := c.Query("id") // shortcut for c.Request.URL.Query().Get("id")
		page := c.DefaultQuery("page", "0")
		hello := json.Hello
		res := fmt.Sprintf("id: %s; page: %s; hello: %s", id, page, hello)
		fmt.Printf(res)

		c.String(http.StatusOK, res)
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}