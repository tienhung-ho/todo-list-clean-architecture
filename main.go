package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-list/common"
	"todo-list/middlewares"
	ginitem "todo-list/modules/items/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database URL from environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Environment variable DATABASE_URL is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Database connection successful")
	fmt.Println(db)

	r := gin.Default()

	// CRUD: Create, Read, Update, Delete
	// /v1/
	r.Use(middlewares.Recovery())
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.GET("/v1/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	})
	r.Run(os.Getenv("PORT")) // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}
