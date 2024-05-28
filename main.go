package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// "id" int8 NOT NULL DEFAULT nextval('todo_items_id_seq'::regclass),
//     "title" varchar(255) NOT NULL,
//     "image" json,
//     "description" text,
//     "status" "public"."status_enum" DEFAULT 'Doing'::status_enum,
//     "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
