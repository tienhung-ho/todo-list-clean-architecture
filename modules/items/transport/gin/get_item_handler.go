package ginitem

import (
	"net/http"
	"strconv"
	"todo-list/common"
	"todo-list/modules/items/business"
	"todo-list/modules/items/storage/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// [GET], /v1/items/:id
func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := postgres.NewPostgresStore(db)
		biz := business.NewGetItemBiz(store)

		data, err := biz.GetItemById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(data))

	}
}
