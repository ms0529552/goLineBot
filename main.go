package main

import (
	"github.com/gin-gonic/gin"
	// 	"context"
	//     "time"
	//     "go.mongodb.org/mongo-driver/mongo"make
	//     "go.mongodb.org/mongo-driver/mongo/options"
	//     "go.mongodb.org/mongo-driver/mongo/readpref"
	// 	"github.com/spf13/cobra"
)

func main() {
	app := gin.Default()
	app.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
