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

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := postgres.NewPostgresStore(db)
		biz := business.NewDeleteItemBiz(store)

		if err := biz.DeleteItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(true))

	}
}
