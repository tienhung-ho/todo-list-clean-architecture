package middlewares

import (
	"net/http"
	"todo-list/common"

	"github.com/gin-gonic/gin"
)

func Recovery() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))
				}

				panic(r)
			}
		}()
		c.Next()
	}
}
