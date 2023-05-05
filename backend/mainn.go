package main

import (
	"backend/src/database"

	"backend/src"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	r.Use(cors.Default())

	// r.GET("/api/products", src.Index)
	r.GET("/api/product/*pertanyaan", src.Show)
	r.POST("/api/algooption", src.AlgoOption)
	// r.POST("/api/product", src.Create)
	// r.PUT("/api/product/:id", src.Update)
	// r.DELETE("/api/product", src.Delete)

	// r.GET("/", func(c *gin.Context) {
	// 	c.File("../frontend/src/tes.html")
	// })

	r.Run(":8000")
}
