package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Post struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

var db = []Post{{Name: "thing", Content: "well yes"}, {Name: "thing2", Content: "well yes2"}, {Name: "thing3", Content: "well yes3"}}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/posts", func(c *gin.Context) {
		var values = db

		c.JSON(http.StatusOK, values)
	})

	r.GET("/post/:id", func(c *gin.Context) {
		var id, _ = strconv.ParseUint(c.Params.ByName("id"), 10, 1)

		var value = db[id]

		c.JSON(http.StatusOK, value)
	})

	r.POST("/post", func(c *gin.Context) {
		var post = Post{}

		if err := c.BindJSON(&post); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)

			return
		}

		db = append(db, post)

		c.JSON(http.StatusOK, gin.H{"message": "Post created."})
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080

	r.Run(":8080")
}
